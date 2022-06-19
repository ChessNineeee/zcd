package client

import (
	"github.com/desertbit/grumble"
	"time"
	"zcd/core"
)

var uploadTLSConnQueue chan *TLSConn
var uploadImmediateQueue chan *TLSConn
var uploadMetaQueue chan *TLSConn
var uploadRequestQueue chan *TLSConn

func init() {
	uploadTLSConnQueue = make(chan *TLSConn, 1024)
	uploadImmediateQueue = make(chan *TLSConn, 1024)
	uploadMetaQueue = make(chan *TLSConn, 1024)
	uploadRequestQueue = make(chan *TLSConn, 1024)
}

func immediateRequestHandler(queue chan *TLSConn) error {
	conn := <-queue
	err := sendRequest(conn, createUploadImmediateRequest(conn.Context))
	if err != nil {
		return err
	}

	err = receiveResponse(conn, createUploadCheckResponse(conn.Context))
	if err != nil {
		return err
	}

	uploadMetaQueue <- conn
	return nil
}

func sendRequest(conn *TLSConn, able core.IEncodingAble) error {
	bytes, err := able.WireEncode()
	if err != nil {
		return err
	}

	_, err = conn.Write(bytes)
	return err
}

func receiveResponse(conn *TLSConn, able core.IEncodingAble) error {
	bytes := make([]byte, 0)
	_, err := conn.Read(bytes)
	if err != nil {
		return err
	}
	err = able.WireDecode(bytes)
	if err != nil {
		return err
	}
	return nil
}

func createUploadImmediateRequest(c *grumble.Context) *core.RequestImmediate {
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

func createUploadMetaRequest(c *grumble.Context, fileSize uint64, nameHash uint32) *core.RequestUploadMeta {
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

func createUploadWorkerRequest(c *grumble.Context, nameHash, workerId uint32, content []byte) *core.RequestUploadWorker {
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
