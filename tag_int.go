package nbt

import (
	"encoding/binary"
	"io"
)

// NewIntTag returns a new Int tag
func NewIntTag(name string, val int32) *Int {
	return &Int{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Int is a tag for an int
type Int struct {
	*tagBase
	Value int32
}

// ID returns tag id
func (t *Int) ID() ID {
	return IDTagInt
}

func (t *Int) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readUint32(reader, order)
	if err != nil {
		return err
	}
	t.Value = int32(val)
	return nil
}

func (t *Int) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeUint32(writer, order, uint32(t.Value))
}

// ToByte returns value as byte
func (t *Int) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToUInt returns value as uint
func (t *Int) ToUint() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Int) ToUint8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Int) ToUint16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Int) ToUint32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Int) ToUint64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Int) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Int) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}
