package nbt

import (
	"encoding/binary"
	"io"
)

// NewFloatTag returns a new Float tag
func NewFloatTag(name string, val float32) *Float {
	return &Float{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Float is a tag for a float
type Float struct {
	*tagBase
	Value float32
}

// ID returns tag id
func (t *Float) ID() ID {
	return IDTagFloat
}

func (t *Float) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readFloat32(reader, order)
	if err != nil {
		return err
	}
	t.Value = val
	return nil
}

func (t *Float) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeFloat32(writer, order, t.Value)
}

// ToFloat32 returns value as float32
func (t *Float) ToFloat32() (float32, error) {
	return t.Value, nil
}

// ToFloat64 returns value as float64
func (t *Float) ToFloat64() (float64, error) {
	return float64(t.Value), nil
}
