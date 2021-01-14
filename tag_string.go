package nbt

import (
	"encoding/binary"
	"errors"
	"io"
	"strconv"
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

// ID returns tag id
func (t *String) ID() ID {
	return IDTagString
}

func (t *String) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	val, err := readString(reader, order)
	if err != nil {
		return err
	}
	t.Value = val
	return nil
}

func (t *String) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	return writeString(writer, order, t.Value)
}

// ToBool returns value as bool
func (t *String) ToBool() (bool, error) {
	return len(t.Value) > 0, nil
}

// ToByte returns value as byte
func (t *String) ToByte() (byte, error) {
	v, err := strconv.Atoi(t.Value)
	return byte(v), err
}

// ToRune returns value as rune
func (t *String) ToRune() (rune, error) {
	v, err := strconv.Atoi(t.Value)
	return rune(v), err
}

// ToInt returns value as int
func (t *String) ToInt() (int, error) {
	v, err := strconv.Atoi(t.Value)
	return v, err
}

// ToUInt returns value as uint
func (t *String) ToUint() (uint, error) {
	// Thanks: licensed CC BY 3.0 by korthaj
	// from https://yourbasic.org/golang/max-min-int-uint/
	v, err := strconv.ParseUint(t.Value, 10, 32<<(^uint(0)>>63))
	return uint(v), err
}

// ToUInt8 returns value as uint8
func (t *String) ToUInt8() (uint8, error) {
	v, err := strconv.ParseUint(t.Value, 10, 8)
	return uint8(v), err
}

// ToUInt16 returns value as uint16
func (t *String) ToUInt16() (uint16, error) {
	v, err := strconv.ParseUint(t.Value, 10, 16)
	return uint16(v), err
}

// ToUInt32 returns value as uint32
func (t *String) ToUInt32() (uint32, error) {
	v, err := strconv.ParseUint(t.Value, 10, 32)
	return uint32(v), err
}

// ToUInt64 returns value as uint64
func (t *String) ToUInt64() (uint64, error) {
	return strconv.ParseUint(t.Value, 10, 64)
}

// ToInt8 returns value as int8
func (t *String) ToInt8() (int8, error) {
	v, err := strconv.Atoi(t.Value)
	return int8(v), err
}

// ToInt16 returns value as int16
func (t *String) ToInt16() (int16, error) {
	v, err := strconv.Atoi(t.Value)
	return int16(v), err
}

// ToInt32 returns value as int32
func (t *String) ToInt32() (int32, error) {
	v, err := strconv.Atoi(t.Value)
	return int32(v), err
}

// ToInt64 returns value as int64
func (t *String) ToInt64() (int64, error) {
	return strconv.ParseInt(t.Value, 10, 64)
}

// ToFloat32 returns value as float32
func (t *String) ToFloat32() (float32, error) {
	v, err := strconv.ParseFloat(t.Value, 32)
	return float32(v), err
}

// ToFloat64 returns value as float64
func (t *String) ToFloat64() (float64, error) {
	return strconv.ParseFloat(t.Value, 64)
}

// ToByteArray returns value as []byte
func (t *String) ToByteArray() ([]byte, error) {
	return []byte(t.Value), nil
}

// ToString returns value as string
func (t *String) ToString() (string, error) {
	return t.Value, nil
}

// ToIntArray returns value as []int32
func (t *String) ToIntArray() ([]int32, error) {
	return nil, errors.New("couldn't cast to []int32")
}

// ToLongArray returns value as []int64
func (t *String) ToLongArray() ([]int64, error) {
	return nil, errors.New("couldn't cast to []int64")
}
