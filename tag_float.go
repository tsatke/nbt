package nbt

import (
	"encoding/binary"
	"io"
)

// NewFloatTag returns a new Float tag.
func NewFloatTag(name string, val float32) *Float {
	return &Float{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Float is a float tag.
type Float struct {
	*tagBase
	Value float32
}

// ID returns this tag's id.
func (t *Float) ID() ID {
	return IDTagFloat
}

// ReadFrom reads a float from the given reader.
func (t *Float) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readFloat32(reader, order)
	if err != nil {
		return err
	}
	t.Value = val
	return nil
}

// WriteTo writes a float to the given writer.
func (t *Float) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeFloat32(writer, order, t.Value)
}
