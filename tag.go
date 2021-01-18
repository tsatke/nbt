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
)

func newTagFromID(id ID) (Tag, error) {
	if id >= NumIDTags {
		return nil, fmt.Errorf("unknown tag %s", id)
	}
	return tagGen[id](), nil
}

// Identifier is something with an ID.
type Identifier interface {
	ID() ID
}

// Namer is something that has a name. The name is mutable.
type Namer interface {
	Name() string
	SetName(string)
}

// ReaderFrom is something that can read from a reader respecting
// a byte order.
type ReaderFrom interface {
	ReadFrom(io.Reader, binary.ByteOrder) error
}

// WriterTo is something that can write to a writer respecting
// a byte order.
type WriterTo interface {
	WriteTo(io.Writer, binary.ByteOrder) error
}

// Tag is an NBT tag.
type Tag interface {
	Identifier
	Namer
	ReaderFrom
	WriterTo
}
