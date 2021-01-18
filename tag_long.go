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

// ID returns this tag's id.
func (t *Long) ID() ID {
	return IDTagLong
}

// ReadFrom reads a long from the given reader.
func (t *Long) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readUint64(reader, order)
	if err != nil {
		return err
	}
	t.Value = int64(val)
	return nil
}

// WriteTo writes this long to the given reader.
func (t *Long) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeUint64(writer, order, uint64(t.Value))
}
