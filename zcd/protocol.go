package zcd

type RequestType uint8

const (
	DUPLICATECHECK RequestType = iota
	UPLOAD
	DELETE
	DOWNLOAD
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
	Length uint32
}

func (r *RequestBase) WireEncode() ([]byte, error) {

}

func (r *RequestType) WireDecode([]byte) error {

}

// RequestImmediate 秒传协议请求结构体
//
type RequestImmediate struct {
	RequestBase
	ContentHash uint32
	NameHash    uint32
}

func (r *RequestImmediate) WireEncode() ([]byte, error) {

}

func (r *RequestImmediate) WireDecode() error {

}

// FileMeta 文件元信息结构体
type FileMeta struct {
	FileSize uint64
	NameHash uint32
	FileName string
}

// RequestUploadMeta 元数据上传请求结构体
//
type RequestUploadMeta struct {
	RequestBase
	RoutinesNum uint32 // 上传使用协程数
	FileMeta
}

func (r *RequestUploadMeta) WireEncode() ([]byte, error) {

}

func (r *RequestUploadMeta) WireDecode() error {

}

// RequestUploadContent 数据分片上传请求结构体
//
type RequestUploadContent struct {
	RequestBase
	NameHash uint32 // NameHash 表明该分片属于哪个文件
	Content  []byte
}

func (r *RequestUploadContent) WireEncode() ([]byte, error) {

}

func (r *RequestUploadContent) WireDecode() error {

}

// RequestDelete 数据删除请求结构体
//
type RequestDelete struct {
	RequestBase
	NameHash uint32 // NameHash 表明删除的文件标识(哈希值)
}

func (r *RequestDelete) WireEncode() ([]byte, error) {

}

func (r *RequestDelete) WireDecode() error {

}

// RequestList 数据查询请求结构体
//
type RequestList struct {
	RequestBase
}

func (r *RequestList) WireEncode() ([]byte, error) {

}

func (r *RequestList) WireDecode() error {

}
