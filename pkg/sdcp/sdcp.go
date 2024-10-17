package sdcp

import (
	"github.com/loopholelabs/logging/types"
	"net"
)

type SDCP struct {
	logger types.Logger
	conn   *net.UDPConn
}

func New(logger types.Logger) (*SDCP, error) {
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, err
	}
	return &SDCP{conn: conn, logger: logger.SubLogger("SDCP")}, nil
}

func (s *SDCP) Close() error {
	return s.conn.Close()
}
