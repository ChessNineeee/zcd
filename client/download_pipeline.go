package client

import (
	"github.com/desertbit/grumble"
	"zcd/core"
)

var downloadTLSConnQueue chan *TLSConn

func init() {
	downloadTLSConnQueue = make(chan *TLSConn, 1024)
}

func createDownloadRequest(c *grumble.Context) *core.RequestDownload {
	r := core.RequestDownload{
		RoutinesNum: core.CommonUint32{
			Value: uint32(c.Flags.Uint64("threads")),
		},
		NameHash: core.CommonUint32{
			Value: uint32(c.Args.Uint64("fileId")),
		},
	}

	r.RequestBase.Length.Value = r.Size()
	return &r
}

func createDownloadWorkerRequest(c *grumble.Context, rId, nameHash uint32) *core.RequestDownloadWorker {
	r := core.RequestDownloadWorker{
		RoutinesId: core.CommonUint32{
			rId,
		},
		NameHash: core.CommonUint32{
			nameHash,
		},
	}
	r.RequestBase.Length.Value = r.Size()

	return &r
}
