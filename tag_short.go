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

// ID returns this tag's id.
func (t *Short) ID() ID {
	return IDTagShort
}

// ReadFrom reads a short from the given reader.
func (t *Short) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readUint16(reader, order)
	if err != nil {
		return err
	}
	t.Value = int16(val)
	return nil
}

// WriteTo writes this short to the given writer.
func (t *Short) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeUint16(writer, order, uint16(t.Value))
}
