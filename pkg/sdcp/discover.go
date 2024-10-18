package sdcp

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"os"
	"time"

	"github.com/loopholelabs/logging/types"
)

var (
	ErrUnableCreateUDPSocket = errors.New("unable to create udp socket")
	ErrBroadcastFailed       = errors.New("broadcast failed")
	ErrReadingUDPSocket      = errors.New("reading udp socket failed")
)

const (
	BroadcastIP   = "255.255.255.255"
	BroadcastPort = 3000

	maximumDiscoverTime = 5 * time.Second
)

var (
	broadcastAddress = &net.UDPAddr{
		IP:   net.ParseIP(BroadcastIP),
		Port: BroadcastPort,
	}
	discoverMessage = []byte("M99999")
)

func Discover(logger types.Logger, ctx context.Context) ([]DiscoverMessage, error) {
	connection, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, errors.Join(ErrUnableCreateUDPSocket, err)
	}

	l := logger.SubLogger("discover")
	l.Debug().Str("listen", connection.LocalAddr().String()).Msg("broadcasting discover message")
	_, err = connection.WriteToUDP(discoverMessage, broadcastAddress)
	if err != nil {
		_ = connection.Close()
		return nil, errors.Join(ErrBroadcastFailed, err)
	}

	_ctx, cancel := context.WithTimeout(ctx, maximumDiscoverTime)
	defer cancel()

	var discovered []DiscoverMessage
	buffer := make([]byte, 8192)
	var n int

	for {
		select {
		case <-_ctx.Done():
			_ = connection.Close()
			return discovered, nil
		default:
			err = connection.SetReadDeadline(time.Now().Add(timeout))
			if err != nil {
				_ = connection.Close()
				return nil, errors.Join(ErrReadingUDPSocket, err)
			}
			n, _, err = connection.ReadFromUDP(buffer)
			if err != nil {
				if errors.Is(err, os.ErrDeadlineExceeded) {
					continue
				}
				_ = connection.Close()
				return nil, errors.Join(ErrReadingUDPSocket, err)
			}
			var message DiscoverMessage
			err = json.Unmarshal(buffer[:n], &message)
			if err != nil {
				_ = connection.Close()
				return nil, errors.Join(ErrReadingUDPSocket, err)
			}
			l.Debug().Str("id", message.ID).Str("machine", message.Data.MainboardID).Str("IP", message.Data.MainboardIP).Msg("discovered device")
			discovered = append(discovered, message)
		}
	}
}
