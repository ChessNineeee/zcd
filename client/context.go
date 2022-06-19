package client

import (
	"crypto/tls"
	"github.com/desertbit/grumble"
)

type TLSConn struct {
	*tls.Conn
	*grumble.Context
}
