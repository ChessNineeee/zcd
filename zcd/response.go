package zcd

const (
	UPLOADCHECK RequestType = iota
	UPLOADRESULT
	DELETERESULT
	LISTRESULT
	DOWNLOADCHECK
	DOWNLOADRESULT
)

type ResponseBase struct {
	Length CommonUint32
}

func (r *ResponseBase) WireEncode() ([]byte, error) {
	return r.Length.WireEncode()
}

func (r *ResponseBase) WireDecode(bytes []byte) error {
	return r.Length.WireDecode(bytes)
}

type ResponseUploadCheck struct {
	ResponseBase
	CommonBool
}

func (r *ResponseUploadCheck) WireEncode() ([]byte, error) {
	res, err := r.Length.WireEncode()
	if err != nil {
		return nil, err
	}

	boolRes, err := r.CommonBool.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, boolRes...)
	return res, nil
}

func (r *ResponseUploadCheck) WireDecode(bytes []byte) error {
	index := 0

	err := r.Length.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = r.CommonBool.WireDecode(bytes[index:r.Length.Value])
	return err
}

type ResponseUploadResult struct {
	ResponseBase
	CommonBool
}

func (r *ResponseUploadResult) WireEncode() ([]byte, error) {
	res, err := r.Length.WireEncode()
	if err != nil {
		return nil, err
	}

	boolRes, err := r.CommonBool.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, boolRes...)
	return res, nil
}

func (r *ResponseUploadResult) WireDecode(bytes []byte) error {
	index := 0

	err := r.Length.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = r.CommonBool.WireDecode(bytes[index:r.ResponseBase.Length.Value])
	return err
}

type ResponseDeleteResult struct {
	ResponseBase
	CommonBool
}

func (r *ResponseDeleteResult) WireEncode() ([]byte, error) {
	res, err := r.Length.WireEncode()
	if err != nil {
		return nil, err
	}

	boolRes, err := r.CommonBool.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, boolRes...)
	return res, nil
}

func (r *ResponseDeleteResult) WireDecode(bytes []byte) error {
	index := 0

	err := r.Length.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = r.CommonBool.WireDecode(bytes[index:r.Length.Value])
	return err
}

// FileInfo 文件信息结构体
type FileInfo struct {
	TSize    CommonUint32
	NameHash CommonUint32
	FileSize CommonUint64
	Time     CommonUint64
	FileName string
}

func (f *FileInfo) WireEncode() ([]byte, error) {
	res, err := f.TSize.WireEncode()
	if err != nil {
		return nil, err
	}

	nameRes, err := f.NameHash.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, nameRes...)

	sizeRes, err := f.FileSize.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, sizeRes...)

	timeRes, err := f.Time.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, timeRes...)

	res = append(res, []byte(f.FileName)...)

	return res, nil
}

func (f *FileInfo) WireDecode(bytes []byte) error {
	index := 0

	err := f.TSize.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = f.NameHash.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = f.FileSize.WireDecode(bytes[index : index+8])
	if err != nil {
		return err
	}
	index += 8

	err = f.Time.WireDecode(bytes[index : index+8])
	if err != nil {
		return err
	}
	index += 8

	f.FileName = string(bytes[index:f.TSize.Value])
	return nil
}

type ResponseListResult struct {
	ResponseBase
	CommonBool
	FileInfoList []*FileInfo
}

func (r *ResponseListResult) WireEncode() ([]byte, error) {
	res, err := r.ResponseBase.WireEncode()
	if err != nil {
		return nil, err
	}

	boolRes, err := r.CommonBool.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, boolRes...)

	if r.CommonBool.Value {
		for _, f := range r.FileInfoList {
			tRes, err := f.WireEncode()
			if err != nil {
				return nil, err
			}
			res = append(res, tRes...)
		}
	}

	return res, nil
}

func (r *ResponseListResult) WireDecode(bytes []byte) error {
	var index uint32 = 0

	err := r.ResponseBase.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = r.CommonBool.WireDecode(bytes[index : index+1])
	if err != nil {
		return err
	}
	index += 1

	if r.CommonBool.Value {
		r.FileInfoList = make([]*FileInfo, 0)
		for index < r.ResponseBase.Length.Value {
			fileInfo := new(FileInfo)
			err = fileInfo.WireDecode(bytes[index:])
			if err != nil {
				return err
			}

			r.FileInfoList = append(r.FileInfoList, fileInfo)
			index += fileInfo.TSize.Value
		}
	}
	return nil
}

type ResponseDownloadCheck struct {
	ResponseBase
	CommonBool
	CommonUint64
}

func (r *ResponseDownloadCheck) WireEncode() ([]byte, error) {
	res, err := r.Length.WireEncode()
	if err != nil {
		return nil, err
	}

	boolRes, err := r.CommonBool.WireEncode()
	if err != nil {
		return nil, err
	}
	res = append(res, boolRes...)

	if r.CommonBool.Value {
		sizeRes, err := r.CommonUint64.WireEncode()
		if err != nil {
			return nil, err
		}
		res = append(res, sizeRes...)
	}

	return res, nil
}

func (r *ResponseDownloadCheck) WireDecode(bytes []byte) error {
	index := 0

	err := r.Length.WireDecode(bytes[index : index+4])
	if err != nil {
		return err
	}
	index += 4

	err = r.CommonBool.WireDecode(bytes[index : index+1])
	if err != nil {
		return err
	}
	index += 1

	if r.CommonBool.Value {
		err = r.CommonUint64.WireDecode(bytes[index:r.ResponseBase.Length.Value])
		if err != nil {
			return err
		}
	}

	return nil
}
