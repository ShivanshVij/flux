package sdcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/loopholelabs/logging/types"
)

var (
	ErrDialFailed               = errors.New("dial failed")
	ErrStatusRefreshFailed      = errors.New("status refresh failed")
	ErrAttributesRefreshFailed  = errors.New("attributes refresh failed")
	ErrEnableDisableVideoFailed = errors.New("enable/disable video failed")
)

const (
	timeout     = 100 * time.Millisecond
	refreshTime = 15 * time.Second
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

	attributesMu   sync.RWMutex
	attributesCond *sync.Cond
	attributes     Attributes

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	requestTopic    string
	responseTopic   string
	statusTopic     string
	attributesTopic string
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
		inflight:        make(map[string]*inflight),
		requestTopic:    fmt.Sprintf("sdcp/request/%s", id),
		responseTopic:   fmt.Sprintf("sdcp/response/%s", id),
		statusTopic:     fmt.Sprintf("sdcp/status/%s", id),
		attributesTopic: fmt.Sprintf("sdcp/attributes/%s", id),
	}

	m.statusCond = sync.NewCond(&m.statusMu)
	m.attributesCond = sync.NewCond(&m.attributesMu)

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

	_, err = m.AttributesRefreshWait(m.ctx)
	if err != nil {
		m.stop()
		m.logger.Error().Err(err).Msg("failed to refresh attributes")
		return nil, errors.Join(ErrAttributesRefreshFailed)
	}

	m.wg.Add(1)
	go m.refresh()

	return m, nil
}

func (m *Machine) StatusRefresh(ctx context.Context) (*StatusRefreshResponse, error) {
	response, err := request(m, CommandStatusRefresh, StatusRefreshRequest{}, ctx)
	if err != nil {
		m.logger.Error().Err(err).Msg("error during status refresh request")
		return nil, errors.Join(ErrStatusRefreshFailed, err)
	}

	var s StatusRefreshResponse
	data, err := json.Marshal(response.Data.Data)
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

func (m *Machine) AttributesRefresh(ctx context.Context) (*AttributesRefreshResponse, error) {
	response, err := request(m, CommandAttributesRefresh, AttributesRefreshRequest{}, ctx)
	if err != nil {
		m.logger.Error().Err(err).Msg("error during attributes refresh request")
		return nil, errors.Join(ErrAttributesRefreshFailed, err)
	}

	var a AttributesRefreshResponse
	data, err := json.Marshal(response.Data.Data)
	if err != nil {
		m.logger.Error().Err(err).Msg("error encoding attributes refresh response")
		return nil, errors.Join(ErrAttributesRefreshFailed, err)
	}
	err = json.Unmarshal(data, &a)
	if err != nil {
		m.logger.Error().Err(err).Msg("error decoding attributes refresh response")
		return nil, errors.Join(ErrAttributesRefreshFailed, err)
	}
	return &a, nil
}

func (m *Machine) AttributesRefreshWait(ctx context.Context) (*Attributes, error) {
	m.attributesMu.Lock()
	_, err := m.AttributesRefresh(ctx)
	if err != nil {
		m.attributesMu.Unlock()
		return nil, err
	}
	m.attributesCond.Wait()
	a := m.attributes
	m.attributesMu.Unlock()
	return &a, nil
}

func (m *Machine) Attributes() *Attributes {
	m.attributesMu.RLock()
	a := m.attributes
	m.attributesMu.RUnlock()
	return &a
}

func (m *Machine) EnableDisableVideo(ctx context.Context, enable bool) (*EnableDisableVideoStreamResponse, error) {
	_enable := EnableDisableDisable
	if enable {
		_enable = EnableDisableEnable
	}

	response, err := request(m, CommandEnableDisableVideoStream, EnableDisableVideoStreamRequest{Enable: _enable}, ctx)
	if err != nil {
		m.logger.Error().Err(err).Msg("error during status refresh request")
		return nil, errors.Join(ErrStatusRefreshFailed, err)
	}

	var a EnableDisableVideoStreamResponse
	data, err := json.Marshal(response.Data.Data)
	if err != nil {
		m.logger.Error().Err(err).Msg("error encoding enable/disable video stream response")
		return nil, errors.Join(ErrEnableDisableVideoFailed, err)
	}
	err = json.Unmarshal(data, &a)
	if err != nil {
		m.logger.Error().Err(err).Msg("error decoding enable/disable video stream response")
		return nil, errors.Join(ErrEnableDisableVideoFailed, err)
	}
	return &a, nil
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
				var status StatusMessage
				err = json.Unmarshal(message, &status)
				if err != nil {
					m.logger.Error().Err(err).Msg("error decoding status message")
					continue
				}
				m.statusMu.Lock()
				m.status = status.Status
				m.statusCond.Broadcast()
				m.statusMu.Unlock()
				m.logger.Debug().Msgf("received status update")
			case m.attributesTopic:
				var attributes AttributesMessage
				err = json.Unmarshal(message, &attributes)
				if err != nil {
					m.logger.Error().Err(err).Msg("error decoding attributes message")
					continue
				}
				m.attributesMu.Lock()
				m.attributes = attributes.Attributes
				m.attributesCond.Broadcast()
				m.attributesMu.Unlock()
				m.logger.Debug().Msgf("received attributes update")
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
			}
			_, err = m.AttributesRefresh(m.ctx)
			if err != nil {
				m.logger.Error().Err(err).Msg("error refreshing attributes")
			}
		}
	}
}

func request[T any](m *Machine, command Command, request T, ctx context.Context) (*Response[any], error) {
	requestID := uuid.New().String()
	msg := &Request[T]{
		TopicMessage: TopicMessage{
			Topic: m.requestTopic,
		},
		Id: identifier,
		Data: RequestData[T]{
			Cmd:         command,
			Data:        request,
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
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-m.ctx.Done():
		return nil, m.ctx.Err()
	case <-i.signal:
	}

	return i.response, nil
}
