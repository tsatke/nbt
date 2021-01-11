package nbt

import (
	"encoding/binary"
	"io"
)

func NewEndTag() *End {
	return &End{
		tagBase: &tagBase{
			name: "end",
		},
	}
}

// End is a end tag
// It shows stream end of tag data
type End struct {
	*tagBase
}

// ID returns tag id
func (t *End) ID() ID {
	return IDTagEnd
}

func (t *End) ReadFrom(_ io.Reader, _ binary.ByteOrder) error {
	return nil
}

func (t *End) WriteTo(w io.Writer, order binary.ByteOrder) error {
	return writeByte(w, order, byte(t.ID()))
}

// ToBool returns value as bool
func (t *End) ToBool() (bool, error) {
	return false, nil
}

// ToByte returns value as byte
func (t *End) ToByte() (byte, error) {
	return 0, nil
}

// ToRune returns value as rune
func (t *End) ToRune() (rune, error) {
	return 0, nil
}

// ToInt returns value as int
func (t *End) ToInt() (int, error) {
	return 0, nil
}

// ToUInt returns value as uint
func (t *End) ToUint() (uint, error) {
	return 0, nil
}

// ToUInt8 returns value as uint8
func (t *End) ToUInt8() (uint8, error) {
	return 0, nil
}

// ToUInt16 returns value as uint16
func (t *End) ToUInt16() (uint16, error) {
	return 0, nil
}

// ToUInt32 returns value as uint32
func (t *End) ToUInt32() (uint32, error) {
	return 0, nil
}

// ToUInt64 returns value as uint64
func (t *End) ToUInt64() (uint64, error) {
	return 0, nil
}

// ToInt8 returns value as int8
func (t *End) ToInt8() (int8, error) {
	return 0, nil
}

// ToInt16 returns value as int16
func (t *End) ToInt16() (int16, error) {
	return 0, nil
}

// ToInt32 returns value as int32
func (t *End) ToInt32() (int32, error) {
	return 0, nil
}

// ToInt64 returns value as int64
func (t *End) ToInt64() (int64, error) {
	return 0, nil
}

// ToFloat32 returns value as float32
func (t *End) ToFloat32() (float32, error) {
	return 0, nil
}

// ToFloat64 returns value as float64
func (t *End) ToFloat64() (float64, error) {
	return 0, nil
}

// ToByteArray returns value as []byte
func (t *End) ToByteArray() ([]byte, error) {
	return []byte{}, nil
}

// ToString returns value as string
func (t *End) ToString() (string, error) {
	return "", nil
}

// ToIntArray returns value as []int32
func (t *End) ToIntArray() ([]int32, error) {
	return []int32{}, nil
}

// ToLongArray returns value as []int64
func (t *End) ToLongArray() ([]int64, error) {
	return []int64{}, nil
}
