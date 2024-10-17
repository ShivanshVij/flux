package sdcp

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"os"
	"time"
)

var (
	ErrUnableCreateUDPSocket = errors.New("unable to create udp socket")
	ErrBroadcastFailed       = errors.New("broadcast failed")
	ErrReadingUDPSocket      = errors.New("reading udp socket failed")
)

const (
	BroadcastIP   = "255.255.255.255"
	BroadcastPort = 3000

	timeout              = 100 * time.Millisecond
	maximumDiscoveryTime = 5 * time.Second
)

var (
	broadcastAddress = &net.UDPAddr{
		IP:   net.ParseIP(BroadcastIP),
		Port: BroadcastPort,
	}
	discoveryMessage = []byte("M99999")
)

func Discover(ctx context.Context) ([]DiscoveryMessage, error) {
	connection, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, errors.Join(ErrUnableCreateUDPSocket, err)
	}

	_, err = connection.WriteToUDP(discoveryMessage, broadcastAddress)
	if err != nil {
		_ = connection.Close()
		return nil, errors.Join(ErrBroadcastFailed, err)
	}

	_ctx, cancel := context.WithTimeout(ctx, maximumDiscoveryTime)
	defer cancel()

	var discovered []DiscoveryMessage
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
			var message DiscoveryMessage
			err = json.Unmarshal(buffer[:n], &message)
			if err != nil {
				_ = connection.Close()
				return nil, errors.Join(ErrReadingUDPSocket, err)
			}
			discovered = append(discovered, message)
		}
	}
}
