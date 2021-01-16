package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

var (
	tagGen = [NumIDTags]func() Tag{
		func() Tag {
			return &End{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &Byte{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &Short{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &Int{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &Long{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &Float{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &Double{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &ByteArray{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &String{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &List{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &Compound{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &IntArray{
				tagBase: &tagBase{},
			}
		},
		func() Tag {
			return &LongArray{
				tagBase: &tagBase{},
			}
		},
	}
	tagNames = [NumIDTags]string{
		"End",
		"Byte",
		"Short",
		"Int",
		"Long",
		"Float",
		"Double",
		"ByteArray",
		"String",
		"List",
		"Compound",
		"IntArray",
		"LongArray",
	}
)

func newTagFromID(id ID) (Tag, error) {
	if id >= NumIDTags {
		return nil, fmt.Errorf("unknown tag %s", id)
	}
	return tagGen[id](), nil
}

type Identifier interface {
	ID() ID
}

type Namer interface {
	Name() string
	SetName(string)
}

type ReaderFrom interface {
	ReadFrom(io.Reader, binary.ByteOrder) error
}

type WriterTo interface {
	WriteTo(io.Writer, binary.ByteOrder) error
}

type NumericConverter interface {
	ToByte() (byte, error)
	ToUint() (uint, error)
	ToUint8() (uint8, error)
	ToUint16() (uint16, error)
	ToUint32() (uint32, error)
	ToUint64() (uint64, error)
}

type FloatingPointConverter interface {
	ToFloat32() (float32, error)
	ToFloat64() (float64, error)
}

type BoolConverter interface {
	ToBool() (bool, error)
}

type RuneConverter interface {
	ToRune() (rune, error)
}

type ArrayConverter interface {
	ToByteArray() ([]byte, error)
	ToIntArray() ([]int32, error)
	ToLongArray() ([]int64, error)
}

type StringConverter interface {
	ToString() (string, error)
}

// Tag is a nbt tag interface
type Tag interface {
	Identifier
	Namer
	ReaderFrom
	WriterTo
}
