package nbt

import (
	"encoding/binary"
	"io"
)

var _ Tag = (*String)(nil)

// NewStringTag returns a new String tag
func NewStringTag(name string, val string) *String {
	return &String{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// String is a tag for string
type String struct {
	*tagBase
	Value string
}

// ID returns this tag's id.
func (t *String) ID() ID {
	return IDTagString
}

// ReadFrom reads a string from the given reader.
func (t *String) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readString(reader, order)
	if err != nil {
		return err
	}
	t.Value = val
	return nil
}

// WriteTo writes this string to the given reader.
func (t *String) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeString(writer, order, t.Value)
}
