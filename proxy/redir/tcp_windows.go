package redir

import (
	"errors"
	"net"

	"github.com/fengyie007/clash/component/socks5"
)

func parserPacket(conn net.Conn) (socks5.Addr, error) {
	return nil, errors.New("Windows not support yet")
}
