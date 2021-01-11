package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

// NewByteArrayTag returns a new ByteArray tag
func NewByteArrayTag(name string, val []byte) *ByteArray {
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
	Value []byte
}

// ID returns tag id
func (t *ByteArray) ID() ID {
	return IDTagByteArray
}

func (t *ByteArray) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	arrLen, err := readUint32(reader, order)
	if err != nil {
		return err
	}

	buf := make([]byte, arrLen)
	if err := read(reader, buf); err != nil {
		return err
	}
	t.Value = buf
	return nil
}

func (t *ByteArray) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	if err := writeUint32(writer, order, uint32(len(t.Value))); err != nil {
		return fmt.Errorf("write length: %w", err)
	}
	if err := write(writer, t.Value); err != nil {
		return err
	}
	return nil
}

// ToByteArray returns value as []byte
func (t *ByteArray) ToByteArray() ([]byte, error) {
	return t.Value, nil
}
