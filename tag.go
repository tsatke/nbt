package nbt

import (
	"encoding/binary"
	"io"
)

var (
	tagGen = [NumIDTags]func() Tag{
		func() Tag { return new(End) },
		func() Tag { return new(Byte) },
		func() Tag { return new(Short) },
		func() Tag { return new(Int) },
		func() Tag { return new(Long) },
		func() Tag { return new(Float) },
		func() Tag { return new(Double) },
		func() Tag { return new(ByteArray) },
		func() Tag { return new(String) },
		func() Tag { return new(List) },
		func() Tag { return new(Compound) },
		func() Tag { return new(IntArray) },
		func() Tag { return new(LongArray) },
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

func newTagFromID(id ID) Tag {
	return tagGen[id]()
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
