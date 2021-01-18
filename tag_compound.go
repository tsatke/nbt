package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

// NewCompoundTag returns a new Compound tag.
func NewCompoundTag(name string, val []Tag) *Compound {
	compound := &Compound{
		tagBase: &tagBase{
			name: name,
		},
		Value: make(map[string]Tag),
	}

	for _, tag := range val {
		compound.Value[tag.Name()] = tag
	}

	return compound
}

// Compound is a list of named tags.
type Compound struct {
	*tagBase
	Value map[string]Tag
}

// ID returns this tag's id.
func (t *Compound) ID() ID {
	return IDTagCompound
}

// ReadFrom reads a compound tag from the given reader.
func (t *Compound) ReadFrom(reader io.Reader, order binary.ByteOrder) error {
	t.Value = make(map[string]Tag)
	decoder := NewDecoder(reader, order)

	for {
		tag, err := decoder.ReadTag()
		if err != nil {
			return fmt.Errorf("read tag: %w", err)
		}
		if tag.ID() == IDTagEnd {
			break
		}

		t.Value[tag.Name()] = tag
	}
	return nil
}

// WriteTo writes this tag to the given writer.
func (t *Compound) WriteTo(writer io.Writer, order binary.ByteOrder) error {
	encoder := NewEncoder(writer, order)
	for _, tag := range t.Value {
		if err := encoder.WriteTag(tag); err != nil {
			return fmt.Errorf("write tag: %w", err)
		}
	}
	if err := NewEndTag().WriteTo(writer, order); err != nil {
		return fmt.Errorf("write end: %w", err)
	}
	return nil
}

// Get returns the named tag with the given name in this compound, or false if no such tag exists.
func (t *Compound) Get(name string) (Tag, bool) {
	tag, ok := t.Value[name]
	return tag, ok
}

// Put puts the given named tag into this compound.
func (t *Compound) Put(tag Tag) {
	t.Value[tag.Name()] = tag
}
