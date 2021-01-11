package nbt

import (
	"encoding/binary"
	"io"
)

// NewByteTag returns a new Byte tag
func NewByteTag(name string, val int8) *Byte {
	return &Byte{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Byte is a tag for a byte
type Byte struct {
	*tagBase
	Value int8
}

// ID returns tag id
func (t *Byte) ID() ID {
	return IDTagByte
}

func (t *Byte) ReadFrom(reader io.Reader, bo binary.ByteOrder) error {
	val, err := readByte(reader, bo)
	if err != nil {
		return err
	}
	t.Value = int8(val)
	return nil
}

func (t *Byte) WriteTo(writer io.Writer, bo binary.ByteOrder) error {
	return writeByte(writer, bo, byte(t.Value))
}

// ToByte returns value as byte
func (t *Byte) ToByte() (byte, error) {
	return byte(t.Value), nil
}

// ToRune returns value as rune
func (t *Byte) ToRune() (rune, error) {
	return rune(t.Value), nil
}

// ToInt returns value as int
func (t *Byte) ToInt() (int, error) {
	return int(t.Value), nil
}

func (t *Byte) ToUint() (uint, error) {
	return uint(t.Value), nil
}

func (t *Byte) ToUint8() (uint8, error) {
	return uint8(t.Value), nil
}

func (t *Byte) ToUint16() (uint16, error) {
	return uint16(t.Value), nil
}

func (t *Byte) ToUint32() (uint32, error) {
	return uint32(t.Value), nil
}

func (t *Byte) ToUint64() (uint64, error) {
	return uint64(t.Value), nil
}

func (t *Byte) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

func (t *Byte) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}
