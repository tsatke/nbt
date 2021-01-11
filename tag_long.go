package nbt

import (
	"encoding/binary"
	"io"
)

// NewLongTag returns a new Long tag
func NewLongTag(name string, val int64) *Long {
	return &Long{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Long is a tag for an int64
type Long struct {
	*tagBase
	Value int64
}

// ID returns tag id
func (t *Long) ID() ID {
	return IDTagLong
}

func (t *Long) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readUint64(reader, order)
	if err != nil {
		return err
	}
	t.Value = int64(val)
	return nil
}

func (t *Long) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeUint64(writer, order, uint64(t.Value))
}

// ToByte returns value as byte
func (t *Long) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToUInt returns value as uint
func (t *Long) ToUint() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Long) ToUint8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Long) ToUint16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Long) ToUint32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Long) ToUint64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Long) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Long) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}
