package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

// NewListTag returns a new List tag.
func NewListTag(name string, val []Tag, typ ID) *List {
	return &List{
		tagBase: &tagBase{
			name: name,
		},
		Value:    val,
		ListType: typ,
	}
}

// List is a list of unnamed tags.
type List struct {
	*tagBase
	Value []Tag

	ListType ID
}

// ID returns this tag's id.
func (t *List) ID() ID {
	return IDTagList
}

// ReadFrom reads a list from the given reader.
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

	t.Value = make([]Tag, 0) // MSER-11 remove preallocation
	for i := uint32(0); i < listLen; i++ {
		tag, err := newTagFromID(t.ListType)
		if err != nil {
			return fmt.Errorf("new tag: %w", err)
		}
		if err := tag.ReadFrom(reader, order); err != nil {
			return fmt.Errorf("read tag: %w", err)
		}
		t.Value = append(t.Value, tag)
	}

	return nil
}

// WriteTo writes this tag to the given writer.
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
