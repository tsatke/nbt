package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

// NewLongArrayTag returns a new LongArray tag
func NewLongArrayTag(name string, val []int64) *LongArray {
	return &LongArray{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// LongArray is a tag for a array of longs
type LongArray struct {
	*tagBase
	Value []int64
}

// ID returns the tag id.
func (t *LongArray) ID() ID {
	return IDTagLongArray
}

// ReadFrom reads a long array from the given reader.
func (t *LongArray) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	arrLen, err := readUint32(reader, order)
	if err != nil {
		return err
	}

	buf := make([]int64, arrLen)
	for i := range buf {
		val, err := readUint64(reader, order)
		if err != nil {
			return fmt.Errorf("element %d: %w", i, err)
		}
		buf[i] = int64(val)
	}
	t.Value = buf
	return nil
}

// WriteTo writes this long array to the given writer.
func (t *LongArray) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	if err := writeUint32(writer, order, uint32(len(t.Value))); err != nil {
		return fmt.Errorf("write length: %w", err)
	}
	for i, val := range t.Value {
		if err := writeUint64(writer, order, uint64(val)); err != nil {
			return fmt.Errorf("element %d: %w", i, err)
		}
	}
	return nil
}
