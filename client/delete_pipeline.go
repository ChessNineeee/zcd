package client

import (
	"github.com/desertbit/grumble"
	"zcd/core"
)

var deleteTLSConnQueue chan *TLSConn

func init() {
	deleteTLSConnQueue = make(chan *TLSConn, 1024)
}

func createDeleteRequest(c grumble.Context) *core.RequestDelete {
	fileId := c.Args.Uint64("fileId")
	r := core.RequestDelete{
		NameHash: core.CommonUint32{
			uint32(fileId),
		},
	}
	r.RequestBase.Length.Value = r.Size()

	return &r
}
