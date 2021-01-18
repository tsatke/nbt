package nbt

import (
	"encoding/binary"
	"io"
)

// NewIntTag returns a new Int tag.
func NewIntTag(name string, val int32) *Int {
	return &Int{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Int is an int tag.
type Int struct {
	*tagBase
	Value int32
}

// ID returns this tag's id.
func (t *Int) ID() ID {
	return IDTagInt
}

// ReadFrom reads an int from the given reader.
func (t *Int) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readUint32(reader, order)
	if err != nil {
		return err
	}
	t.Value = int32(val)
	return nil
}

// WriteTo writes this int to the given writer.
func (t *Int) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeUint32(writer, order, uint32(t.Value))
}
