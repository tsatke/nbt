package nbt

import (
	"encoding/binary"
	"io"
)

// NewDoubleTag returns a new Double tag
func NewDoubleTag(name string, val float64) *Double {
	return &Double{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Double is a tag for float64
type Double struct {
	*tagBase
	Value float64
}

// ID returns tag id
func (t *Double) ID() ID {
	return IDTagDouble
}

func (t *Double) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readFloat64(reader, order)
	if err != nil {
		return err
	}
	t.Value = val
	return nil
}

func (t *Double) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeFloat64(writer, order, t.Value)
}

// ToFloat32 returns value as float32
func (t *Double) ToFloat32() (float32, error) {
	return float32(t.Value), nil
}

// ToFloat64 returns value as float64
func (t *Double) ToFloat64() (float64, error) {
	return t.Value, nil
}
