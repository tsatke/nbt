package nbt

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strconv"
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

// ID returns tag id
func (t *LongArray) ID() ID {
	return IDTagLongArray
}

func (t *LongArray) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	arrLen, err := readUint64(reader, order)
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

func (t *LongArray) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	if err := writeUint64(writer, order, uint64(len(t.Value))); err != nil {
		return fmt.Errorf("write length: %w", err)
	}
	for i, val := range t.Value {
		if err := writeUint64(writer, order, uint64(val)); err != nil {
			return fmt.Errorf("element %d: %w", i, err)
		}
	}
	return nil
}

// ToBool returns value as bool
func (t *LongArray) ToBool() (bool, error) {
	return len(t.Value) > 0, nil
}

// ToByte returns value as byte
func (t *LongArray) ToByte() (byte, error) {
	return 0, errors.New("couldn't cast to byte")
}

// ToRune returns value as rune
func (t *LongArray) ToRune() (rune, error) {
	return 0, errors.New("couldn't cast to rune")
}

// ToInt returns value as int
func (t *LongArray) ToInt() (int, error) {
	return 0, errors.New("couldn't cast to int")
}

// ToUInt returns value as uint
func (t *LongArray) ToUint() (uint, error) {
	return 0, errors.New("couldn't cast to uint")
}

// ToUInt8 returns value as uint8
func (t *LongArray) ToUInt8() (uint8, error) {
	return 0, errors.New("couldn't cast to uint8")
}

// ToUInt16 returns value as uint16
func (t *LongArray) ToUInt16() (uint16, error) {
	return 0, errors.New("couldn't cast to uint16")
}

// ToUInt32 returns value as uint32
func (t *LongArray) ToUInt32() (uint32, error) {
	return 0, errors.New("couldn't cast to uint32")
}

// ToUInt64 returns value as uint64
func (t *LongArray) ToUInt64() (uint64, error) {
	return 0, errors.New("couldn't cast to uint64")
}

// ToInt8 returns value as int8
func (t *LongArray) ToInt8() (int8, error) {
	return 0, errors.New("couldn't cast to int8")
}

// ToInt16 returns value as int16
func (t *LongArray) ToInt16() (int16, error) {
	return 0, errors.New("couldn't cast to int16")
}

// ToInt32 returns value as int32
func (t *LongArray) ToInt32() (int32, error) {
	return 0, errors.New("couldn't cast to int32")
}

// ToInt64 returns value as int64
func (t *LongArray) ToInt64() (int64, error) {
	return 0, errors.New("couldn't cast to int64")
}

// ToFloat32 returns value as float32
func (t *LongArray) ToFloat32() (float32, error) {
	return 0, errors.New("couldn't cast to float32")
}

// ToFloat64 returns value as float64
func (t *LongArray) ToFloat64() (float64, error) {
	return 0, errors.New("couldn't cast to float64")
}

// ToByteArray returns value as []byte
func (t *LongArray) ToByteArray() ([]byte, error) {
	return nil, errors.New("couldn't cast to []byte")
}

// ToString returns value as string
func (t *LongArray) ToString() (string, error) {
	str := "[ "
	for i, v := range t.Value {
		str += strconv.Itoa(int(v))
		if i != (len(t.Value) - 1) { // not last
			str += ", "
		}
	}

	return str + " ]", nil
}

// ToIntArray returns value as []int32
func (t *LongArray) ToIntArray() ([]int32, error) {
	return nil, errors.New("couldn't cast to []int32")
}

// ToLongArray returns value as []int64
func (t *LongArray) ToLongArray() ([]int64, error) {
	return t.Value, nil
}
