package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

// NewByteArrayTag returns a new ByteArray tag
func NewByteArrayTag(name string, val []int8) *ByteArray {
	return &ByteArray{
		tagBase: &tagBase{
			name: name,
		},
		Value: val,
	}
}

// ByteArray is a tag for []byte
type ByteArray struct {
	*tagBase
	Value []int8
}

// ID returns tag id
func (t *ByteArray) ID() ID {
	return IDTagByteArray
}

// ReadFrom reads a byte array from the reader.
func (t *ByteArray) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	arrLen, err := readUint32(reader, order)
	if err != nil {
		return err
	}

	buf := make([]byte, arrLen)
	if err := read(reader, buf); err != nil {
		return err
	}
	t.Value = make([]int8, arrLen)
	for i := range buf {
		t.Value[i] = int8(buf[i])
	}
	return nil
}

// WriteTo writes this byte array to the given writer.
func (t *ByteArray) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	if err := writeUint32(writer, order, uint32(len(t.Value))); err != nil {
		return fmt.Errorf("write length: %w", err)
	}
	buf := make([]byte, len(t.Value))
	for i := range buf {
		buf[i] = byte(t.Value[i])
	}
	if err := write(writer, buf); err != nil {
		return err
	}
	return nil
}
