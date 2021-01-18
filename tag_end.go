package nbt

import (
	"encoding/binary"
	"io"
)

// NewEndTag creates a new End tag.
func NewEndTag() *End {
	return &End{
		tagBase: &tagBase{
			name: "end",
		},
	}
}

// End is an end tag. This has no data.
type End struct {
	*tagBase
}

// ID returns tag id
func (t *End) ID() ID {
	return IDTagEnd
}

// ReadFrom is an effective noop.
func (t *End) ReadFrom(_ io.Reader, _ binary.ByteOrder) error {
	return nil
}

// WriteTo writes this tag onto the given writer, i.e. writes the end tag id byte.
func (t *End) WriteTo(w io.Writer, order binary.ByteOrder) error {
	return writeByte(w, order, byte(t.ID()))
}
