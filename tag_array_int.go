package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

// NewIntArrayTag returns a new IntArray tag
func NewIntArrayTag(name string, val []int32) *IntArray {
	return &IntArray{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// IntArray is a tag for a int array
type IntArray struct {
	*tagBase
	Value []int32
}

// ID returns tag id
func (t *IntArray) ID() ID {
	return IDTagIntArray
}

// ReadFrom reads an int array from the given reader.
func (t *IntArray) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	arrLen, err := readUint32(reader, order)
	if err != nil {
		return err
	}

	buf := make([]int32, arrLen)
	for i := range buf {
		val, err := readUint32(reader, order)
		if err != nil {
			return fmt.Errorf("element %d: %w", i, err)
		}
		buf[i] = int32(val)
	}
	t.Value = buf
	return nil
}

// WriteTo writes this int array to the given writer.
func (t *IntArray) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	if err := writeUint32(writer, order, uint32(len(t.Value))); err != nil {
		return fmt.Errorf("write length: %w", err)
	}
	for i, val := range t.Value {
		if err := writeUint32(writer, order, uint32(val)); err != nil {
			return fmt.Errorf("element %d: %w", i, err)
		}
	}
	return nil
}
