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

// ToIntArray returns value as []int32
func (t *IntArray) ToIntArray() ([]int32, error) {
	return t.Value, nil
}

// ToLongArray returns value as []int64
func (t *IntArray) ToLongArray() ([]int64, error) {
	result := make([]int64, len(t.Value))
	for i := 0; i < len(result); i++ {
		result[i] = int64(t.Value[i])
	}

	return result, nil
}
