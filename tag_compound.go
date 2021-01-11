package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

// NewCompoundTag returns a new Compound tag
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

// Compound is a map contained tags
type Compound struct {
	*tagBase
	Value map[string]Tag
}

// ID returns tag id
func (t *Compound) ID() ID {
	return IDTagCompound
}

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

func (t *Compound) Get(name string) (Tag, bool) {
	tag, ok := t.Value[name]
	return tag, ok
}

func (t *Compound) Put(tag Tag) {
	t.Value[tag.Name()] = tag
}
