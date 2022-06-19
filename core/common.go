package core

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

func (c *CommonUint32) Size() uint32 {
	return 4
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

func (c *CommonUint64) Size() uint32 {
	return 8
}

type CommonBool struct {
	Value bool
}

func (c *CommonBool) WireEncode() ([]byte, error) {
	res := make([]byte, 1)
	if c.Value {
		res[0] = 1
	} else {
		res[0] = 0
	}
	return res, nil
}

func (c *CommonBool) WireDecode(bytes []byte) error {
	if bytes == nil || len(bytes) != 1 {
		return errors.New("CommonBool:Invalid Bytes Array")
	}
	if bytes[0] == 1 {
		c.Value = true
	} else {
		c.Value = false
	}

	return nil
}

func (c *CommonBool) Size() uint32 {
	return 1
}
