package sdcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/xyproto/randomstring"

	"github.com/loopholelabs/logging/types"
)

var (
	ErrDialFailed          = errors.New("dial failed")
	ErrStatusRefreshFailed = errors.New("status refresh failed")
)

const (
	timeout     = 100 * time.Millisecond
	refreshTime = 5 * time.Second
	apiPort     = 3030
	identifier  = "fluxsdcp"
)

type inflight struct {
	signal   chan struct{}
	response *Response[any]
}

type Machine struct {
	logger types.Logger
	id     string
	ip     string

	url  *url.URL
	conn *websocket.Conn

	inflightMu sync.RWMutex
	inflight   map[string]*inflight

	statusMu   sync.RWMutex
	statusCond *sync.Cond
	status     Status

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	responseTopic string
	statusTopic   string
}

func newMachine(id string, ip string, logger types.Logger) (*Machine, error) {
	m := &Machine{
		logger: logger.SubLogger("machine").With().Str("id", id).Str("ip", ip).Logger(),
		id:     id,
		ip:     ip,
		url: &url.URL{
			Scheme: "ws",
			Host:   fmt.Sprintf("%s:%d", ip, apiPort),
			Path:   "/websocket",
		},
		inflight:      make(map[string]*inflight),
		responseTopic: fmt.Sprintf("sdcp/response/%s", id),
		statusTopic:   fmt.Sprintf("sdcp/status/%s", id),
	}

	m.statusCond = sync.NewCond(&m.statusMu)

	var err error
	m.ctx, m.cancel = context.WithCancel(context.Background())
	m.conn, _, err = websocket.DefaultDialer.DialContext(m.ctx, m.url.String(), nil)
	if err != nil {
		m.logger.Error().Err(err).Msg("failed to connect to machine")
		return nil, errors.Join(ErrDialFailed, err)
	}

	m.logger.Info().Msg("connected to machine")
	m.wg.Add(1)
	go m.handle()

	_, err = m.StatusRefreshWait(m.ctx)
	if err != nil {
		m.stop()
		m.logger.Error().Err(err).Msg("failed to refresh status")
		return nil, errors.Join(ErrStatusRefreshFailed)
	}

	m.wg.Add(1)
	go m.refresh()

	return m, nil
}

func (m *Machine) StatusRefresh(ctx context.Context) (*StatusRefreshResponse, error) {
	requestID := randomstring.HumanFriendlyString(8)
	m.logger.Info().Str("id", requestID).Msg("refreshing status")
	msg := &Request[StatusRefreshRequest]{
		TopicMessage: TopicMessage{
			Topic: fmt.Sprintf("sdcp/request/%s", m.id),
		},
		Id: identifier,
		Data: RequestData[StatusRefreshRequest]{
			Cmd:         CommandStatusRefresh,
			Data:        StatusRefreshRequest{},
			RequestID:   requestID,
			MainboardID: m.id,
			TimeStamp:   int(time.Now().Unix()),
			From:        FromPC,
		},
	}

	i := &inflight{
		signal:   make(chan struct{}),
		response: new(Response[any]),
	}
	m.inflightMu.Lock()
	m.inflight[requestID] = i
	m.inflightMu.Unlock()
	defer func() {
		m.inflightMu.Lock()
		delete(m.inflight, requestID)
		m.inflightMu.Unlock()
	}()

	err := m.conn.WriteJSON(msg)
	if err != nil {
		m.logger.Error().Err(err).Msg("error sending status refresh request")
		return nil, errors.Join(ErrStatusRefreshFailed, err)
	}

	select {
	case <-ctx.Done():
		return nil, errors.Join(ErrStatusRefreshFailed, ctx.Err())
	case <-m.ctx.Done():
		return nil, errors.Join(ErrStatusRefreshFailed, m.ctx.Err())
	case <-i.signal:
	}

	var s StatusRefreshResponse
	data, err := json.Marshal(i.response.Data.Data)
	if err != nil {
		m.logger.Error().Err(err).Msg("error encoding status refresh response")
		return nil, errors.Join(ErrStatusRefreshFailed, err)
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		m.logger.Error().Err(err).Msg("error decoding status refresh response")
		return nil, errors.Join(ErrStatusRefreshFailed, err)
	}
	return &s, nil
}

func (m *Machine) StatusRefreshWait(ctx context.Context) (*Status, error) {
	m.statusMu.Lock()
	_, err := m.StatusRefresh(ctx)
	if err != nil {
		m.statusMu.Unlock()
		return nil, err
	}
	m.statusCond.Wait()
	s := m.status
	m.statusMu.Unlock()
	return &s, nil
}

func (m *Machine) Status() *Status {
	m.statusMu.RLock()
	s := m.status
	m.statusMu.RUnlock()
	return &s
}

func (m *Machine) stop() {
	m.cancel()
	_ = m.conn.Close()
	m.wg.Wait()
}

func (m *Machine) handle() {
	defer m.wg.Done()
	var topicMessage TopicMessage
	var err error
	var message []byte
	for {
		select {
		case <-m.ctx.Done():
			return
		default:
			_, message, err = m.conn.ReadMessage()
			if err != nil {
				m.logger.Error().Err(err).Msg("error reading from websocket")
				return
			}
			err = json.Unmarshal(message, &topicMessage)
			if err != nil {
				m.logger.Error().Err(err).Msg("error decoding topic message")
				continue
			}

			m.logger.Debug().Str("topic", topicMessage.Topic).Msg("received message")
			switch topicMessage.Topic {
			case m.responseTopic:
				var response Response[any]
				err = json.Unmarshal(message, &response)
				if err != nil {
					m.logger.Error().Err(err).Msg("error decoding response message")
					continue
				}
				m.inflightMu.RLock()
				i, ok := m.inflight[response.Data.RequestID]
				m.inflightMu.RUnlock()
				if ok {
					i.response = &response
					close(i.signal)
					m.logger.Debug().Str("request", response.Data.RequestID).Msg("received response")
				} else {
					m.logger.Warn().Str("request", response.Data.RequestID).Msg("unknown request")
				}
			case m.statusTopic:
				var status Status
				err = json.Unmarshal(message, &status)
				if err != nil {
					m.logger.Error().Err(err).Msg("error decoding status message")
					continue
				}
				m.statusMu.Lock()
				m.status = status
				m.statusCond.Broadcast()
				m.statusMu.Unlock()
				m.logger.Debug().Msgf("received status update")
			default:
				m.logger.Warn().Str("topic", topicMessage.Topic).Msg("unknown topic")
			}
		}
	}
}

func (m *Machine) refresh() {
	defer m.wg.Done()
	var err error
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-time.After(refreshTime):
			_, err = m.StatusRefresh(m.ctx)
			if err != nil {
				m.logger.Error().Err(err).Msg("error refreshing status")
				continue
			}
		}
	}
}
