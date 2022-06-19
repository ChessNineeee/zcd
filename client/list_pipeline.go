package client

import (
	"github.com/desertbit/grumble"
	"zcd/core"
)

var listTLSConnQueue chan *TLSConn

func init() {
	listTLSConnQueue = make(chan *TLSConn, 1024)
}

func createListRequest(c grumble.Context) *core.RequestList {
	r := core.RequestList{}
	r.RequestBase.Length.Value = r.Size()

	return &r
}
