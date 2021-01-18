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

// ID returns this tag's id.
func (t *Byte) ID() ID {
	return IDTagByte
}

// ReadFrom reads a byte from the given reader.
func (t *Byte) ReadFrom(reader io.Reader, bo binary.ByteOrder) error {
	val, err := readByte(reader, bo)
	if err != nil {
		return err
	}
	t.Value = int8(val)
	return nil
}

// WriteTo writes this byte to the given writer.
func (t *Byte) WriteTo(writer io.Writer, bo binary.ByteOrder) error {
	return writeByte(writer, bo, byte(t.Value))
}
