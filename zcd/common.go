package zcd

import (
	"encoding/binary"
	"errors"
)

// CommonUint32
//
type CommonUint32 struct {
	Value uint32
}

func (c *CommonUint32) WireEncode() ([]byte, error) {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, c.Value)
	return res, nil
}

func (c *CommonUint32) WireDecode(bytes []byte) error {
	if bytes == nil || len(bytes) != 4 {
		return errors.New("CommonUint32:Invalid Bytes Array")
	}
	c.Value = binary.BigEndian.Uint32(bytes)
	return nil
}

// CommonUint64
//
type CommonUint64 struct {
	Value uint64
}

func (c *CommonUint64) WireEncode() ([]byte, error) {
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, c.Value)
	return res, nil
}

func (c *CommonUint64) WireDecode(bytes []byte) error {
	if bytes == nil || len(bytes) != 8 {
		return errors.New("CommonUint64:Invalid Bytes Array")
	}
	c.Value = binary.BigEndian.Uint64(bytes)
	return nil
}
