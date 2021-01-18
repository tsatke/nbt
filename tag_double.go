package nbt

import (
	"encoding/binary"
	"io"
)

// NewDoubleTag returns a new Double tag.
func NewDoubleTag(name string, val float64) *Double {
	return &Double{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// Double is a double tag, which contains a float64 value.
type Double struct {
	*tagBase
	Value float64
}

// ID returns this tag's id.
func (t *Double) ID() ID {
	return IDTagDouble
}

// ReadFrom reads a double from the given reader.
func (t *Double) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readFloat64(reader, order)
	if err != nil {
		return err
	}
	t.Value = val
	return nil
}

// WriteTo writes this double to the writer.
func (t *Double) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeFloat64(writer, order, t.Value)
}
