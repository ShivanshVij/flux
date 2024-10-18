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
	timeout    = 100 * time.Millisecond
	apiPort    = 3030
	identifier = "fluxsdcp"
)

type Machine struct {
	logger types.Logger
	id     string
	ip     string

	url  *url.URL
	conn *websocket.Conn

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup
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
	}

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

	return m, nil
}

func (m *Machine) StatusRefresh() error {
	m.logger.Info().Msg("refreshing status")
	requestID := randomstring.String(8)
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

	err := m.conn.WriteJSON(msg)
	if err != nil {
		m.logger.Error().Err(err).Msg("error sending status refresh request")
		return errors.Join(ErrStatusRefreshFailed, err)
	}

	return nil
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
			m.logger.Info().Str("topic", topicMessage.Topic).Msg("received message")
		}
	}
}
