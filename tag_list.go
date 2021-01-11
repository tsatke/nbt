package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

// NewListTag returns a new List tag
func NewListTag(name string, val []Tag, typ ID) *List {
	return &List{
		tagBase: &tagBase{
			name: name,
		},
		Value:    val,
		ListType: typ,
	}
}

// List is a container for tags
type List struct {
	*tagBase
	Value []Tag

	ListType ID
}

// ID returns tag id
func (t *List) ID() ID {
	return IDTagList
}

func (t *List) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	idByte, err := readByte(reader, order)
	if err != nil {
		return fmt.Errorf("read list type: %w", err)
	}
	t.ListType = ID(idByte)

	listLen, err := readUint32(reader, order)
	if err != nil {
		return fmt.Errorf("read list length: %w", err)
	}

	t.Value = make([]Tag, listLen)
	for i := range t.Value {
		tag := newTagFromID(t.ListType)
		if err := tag.ReadFrom(reader, order); err != nil {
			return fmt.Errorf("read tag: %w", err)
		}
		t.Value[i] = tag
	}

	return nil
}

func (t *List) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	if err := writeByte(writer, order, byte(t.ListType)); err != nil {
		return fmt.Errorf("write list type: %w", err)
	}
	if err := writeUint32(writer, order, uint32(len(t.Value))); err != nil {
		return fmt.Errorf("write list length: %w", err)
	}
	for i, tag := range t.Value {
		if err := tag.WriteTo(writer, order); err != nil {
			return fmt.Errorf("element %d: %w", i, err)
		}
	}
	return nil
}
