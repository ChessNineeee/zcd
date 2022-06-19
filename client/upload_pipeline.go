package client

import (
	"github.com/desertbit/grumble"
	"time"
	"zcd/core"
)

var uploadTLSConnQueue chan *TLSConn

func init() {
	uploadTLSConnQueue = make(chan *TLSConn, 1024)
}

func createUploadImmediateRequest(c grumble.Context) *core.RequestImmediate {
	fileName := c.Args.String("fileName")

	hash := core.MurmurHash3([]byte(fileName), len(fileName), uint32(time.Now().UnixMilli()))
	r := core.RequestImmediate{
		NameHash: core.CommonUint32{
			hash,
		},
	}
	r.RequestBase.Length.Value = r.Size()

	return &r
}

func createUploadMetaRequest(c grumble.Context, fileSize uint64, nameHash uint32) *core.RequestUploadMeta {
	r := core.RequestUploadMeta{
		RoutinesNum: core.CommonUint32{
			Value: uint32(c.Flags.Uint64("threads")),
		},
		FileMeta: core.FileMeta{
			FileSize: core.CommonUint64{
				Value: fileSize,
			},
			NameHash: core.CommonUint32{
				Value: nameHash,
			},
			FileName: c.Args.String("fileName"),
		},
	}

	r.RequestBase.Length.Value = r.Size()
	return &r
}

func createUploadWorkerRequest(c grumble.Context, nameHash, workerId uint32, content []byte) *core.RequestUploadWorker {
	r := core.RequestUploadWorker{
		NameHash: core.CommonUint32{
			Value: nameHash,
		},
		WorkerId: core.CommonUint32{
			Value: workerId,
		},
		Content: content,
	}

	r.RequestBase.Length.Value = r.Size()
	return &r
}
