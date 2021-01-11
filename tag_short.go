package nbt

import (
	"encoding/binary"
	"io"
)

// NewShortTag returns a new Short tag
func NewShortTag(name string, val int16) *Short {
	return &Short{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Short is a tag for a short
type Short struct {
	*tagBase
	Value int16
}

// ID returns tag id
func (t *Short) ID() ID {
	return IDTagShort
}

func (t *Short) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readUint16(reader, order)
	if err != nil {
		return err
	}
	t.Value = int16(val)
	return nil
}

func (t *Short) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeUint16(writer, order, uint16(t.Value))
}

// ToByte returns value as byte
func (t *Short) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToInt returns value as int
func (t *Short) ToInt() (int, error) {
	return int(t.Value), nil
}

// ToUInt returns value as uint
func (t *Short) ToUint() (uint, error) {
	return uint(t.Value), nil
}

// ToUInt8 returns value as uint8
func (t *Short) ToUint8() (uint8, error) {
	return uint8(t.Value), nil
}

// ToUInt16 returns value as uint16
func (t *Short) ToUint16() (uint16, error) {
	return uint16(t.Value), nil
}

// ToUInt32 returns value as uint32
func (t *Short) ToUint32() (uint32, error) {
	return uint32(t.Value), nil
}

// ToUInt64 returns value as uint64
func (t *Short) ToUint64() (uint64, error) {
	return uint64(t.Value), nil
}

// ToFloat32 returns value as float32
func (t *Short) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Short) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}
