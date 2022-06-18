package zcd

import (
	"errors"
)

type RequestType uint8

const (
	DUPLICATECHECK RequestType = iota
	UPLOAD
	UPLOADWORKER
	DELETE
	DOWNLOAD
	DOWNLOADWORKER
)

// IEncodingAble 可编码结构体接口
//
type IEncodingAble interface {
	WireEncode() []byte      // 将结构体线速编码为字节数组
	WireDecode([]byte) error // 将字节数组线速解码为结构体
}

// RequestBase 请求长度结构体
//
type RequestBase struct {
	Length CommonUint32
}

func (r *RequestBase) WireEncode() ([]byte, error) {
	return r.Length.WireEncode()
}

func (r *RequestBase) WireDecode(bytes []byte) error {
	return r.Length.WireDecode(bytes)
}

// RequestImmediate 秒传协议请求结构体
//
type RequestImmediate struct {
	RequestBase
	ContentHash CommonUint32
	NameHash    CommonUint32
}

func (r *RequestImmediate) WireEncode() ([]byte, error) {
	res, err := r.RequestBase.WireEncode()
	if err != nil {
		return nil, err
	}

	hashRes, err := r.ContentHash.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, hashRes...)

	nameRes, err := r.NameHash.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, nameRes...)

	return res, nil
}

func (r *RequestImmediate) WireDecode(bytes []byte) error {
	index := 0
	// RequestBase self decode
	err := r.RequestBase.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	// ContentHash
	err = r.ContentHash.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	// NameHash
	err = r.NameHash.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}

	return nil
}

// FileMeta 文件元信息结构体
type FileMeta struct {
	FileSize CommonUint64
	NameHash CommonUint32
	FileName string
}

func (f *FileMeta) WireEncode() ([]byte, error) {
	res, err := f.FileSize.WireEncode()
	if err != nil {
		return nil, err
	}

	hashRes, err := f.NameHash.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, hashRes...)

	res = append(res, []byte(f.FileName)...)

	return res, nil
}

func (f *FileMeta) WireDecode(bytes []byte) error {
	index := 0

	// FileSize
	err := f.FileSize.WireDecode(bytes[index : index+8])
	if err != nil {
		return err
	}
	index += 8

	// NameHash
	err = f.NameHash.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	// FileName
	f.FileName = string(bytes[index:])

	return nil
}

// RequestUploadMeta 元数据上传请求结构体
//
type RequestUploadMeta struct {
	RequestBase
	RoutinesNum CommonUint32 // 上传使用协程数
	FileMeta
}

func (r *RequestUploadMeta) WireEncode() ([]byte, error) {
	res, err := r.RequestBase.WireEncode()
	if err != nil {
		return nil, err
	}

	numRes, err := r.RoutinesNum.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, numRes...)

	metaRes, err := r.FileMeta.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, metaRes...)

	return res, nil
}

func (r *RequestUploadMeta) WireDecode(bytes []byte) error {
	index := 0

	err := r.RequestBase.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = r.RoutinesNum.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = r.FileMeta.WireDecode(bytes[index:])
	return err
}

// RequestUploadWorker 数据分片上传请求结构体
//
type RequestUploadWorker struct {
	RequestBase
	NameHash CommonUint32 // NameHash 表明该分片属于哪个文件
	WorkerId CommonUint32 // WorkerId
	Content  []byte
}

func (r *RequestUploadWorker) WireEncode() ([]byte, error) {
	res, err := r.RequestBase.WireEncode()
	if err != nil {
		return nil, err
	}

	nameRes, err := r.NameHash.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, nameRes...)

	workerIdRes, err := r.WorkerId.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, workerIdRes...)

	if r.Content == nil {
		return nil, errors.New("RequestUploadWorker: Bad Content")
	}
	res = append(res, r.Content...)

	return res, nil
}

func (r *RequestUploadWorker) WireDecode(bytes []byte) error {
	index := 0

	// RequestBase
	err := r.RequestBase.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	// NameHash
	err = r.NameHash.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	// WorkerId
	err = r.WorkerId.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	// Content
	r.Content = bytes[index:]

	return nil
}

// RequestDelete 数据删除请求结构体
//
type RequestDelete struct {
	RequestBase
	NameHash CommonUint32 // NameHash 表明删除的文件标识(哈希值)
}

func (r *RequestDelete) WireEncode() ([]byte, error) {
	res, err := r.RequestBase.WireEncode()
	if err != nil {
		return nil, err
	}

	nameRes, err := r.NameHash.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, nameRes...)

	return res, nil
}

func (r *RequestDelete) WireDecode(bytes []byte) error {
	index := 0

	// RequestBase
	err := r.RequestBase.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	// NameHash
	err = r.NameHash.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}

	return nil
}

// RequestList 数据查询请求结构体
//
type RequestList struct {
	RequestBase
}

func (r *RequestList) WireEncode() ([]byte, error) {
	res, err := r.RequestBase.WireEncode()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RequestList) WireDecode(bytes []byte) error {
	index := 0

	// RequestBase
	err := r.RequestBase.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}

	return nil
}

type RequestDownload struct {
	RequestBase
	RoutinesNum CommonUint32 // 下载使用协程数
	NameHash    CommonUint32 // 文件标识
}

func (r *RequestDownload) WireEncode() ([]byte, error) {
	res, err := r.RequestBase.WireEncode()
	if err != nil {
		return nil, err
	}

	numRes, err := r.RoutinesNum.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, numRes...)

	nameRes, err := r.NameHash.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, nameRes...)

	return res, nil
}

func (r *RequestDownload) WireDecode(bytes []byte) error {
	index := 0

	// RequestBase
	err := r.RequestBase.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}

	// Routines Num
	err = r.RoutinesNum.WireDecode(bytes[index : index+4])
	index += 4

	// Name Hash
	err = r.NameHash.WireDecode(bytes[index : index+4])

	return nil
}

type RequestDownloadWorker struct {
	RequestBase
	RoutinesId CommonUint32 // 下载使用协程Id
	NameHash   CommonUint32 // 文件标识
}

func (r *RequestDownloadWorker) WireEncode() ([]byte, error) {
	res, err := r.RequestBase.WireEncode()
	if err != nil {
		return nil, err
	}

	rIdRes, err := r.RoutinesId.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, rIdRes...)

	nameRes, err := r.NameHash.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, nameRes...)

	return res, nil
}

func (r *RequestDownloadWorker) WireDecode(bytes []byte) error {
	index := 0

	// RequestBase
	err := r.RequestBase.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}

	// Routines Num
	err = r.RoutinesId.WireDecode(bytes[index : index+4])
	index += 4

	// Name Hash
	err = r.NameHash.WireDecode(bytes[index : index+4])

	return nil
}
