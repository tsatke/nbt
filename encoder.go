package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// Encoder describes a component that is capable of writing
// NBT tags to some sink.
type Encoder interface {
	WriteTag(Tag) error
}

type encoder struct {
	w  io.Writer
	bo binary.ByteOrder
}

// NewEncoder creates a new Encoder that will encode NBT tags
// with the given byte order and write them on the given writer.
func NewEncoder(target io.Writer, byteOrder binary.ByteOrder) Encoder {
	return &encoder{
		w:  target,
		bo: byteOrder,
	}
}

func (e encoder) WriteTag(tag Tag) error {
	if err := writeByte(e.w, e.bo, byte(tag.ID())); err != nil {
		return fmt.Errorf("write ID: %w", err)
	}
	if err := writeString(e.w, e.bo, tag.Name()); err != nil {
		return fmt.Errorf("write tag name: %w", err)
	}
	if err := tag.WriteTo(e.w, e.bo); err != nil {
		return fmt.Errorf("write %s with name '%s': %w", tag.Name(), tag.ID(), err)
	}
	return nil
}

func writeByte(w io.Writer, _ binary.ByteOrder, b byte) error { // anonymous arg to keep signature consistent
	return write(w, []byte{b})
}

func writeUint16(w io.Writer, order binary.ByteOrder, i uint16) error {
	buf := make([]byte, 2)
	order.PutUint16(buf, i)
	return write(w, buf)
}

func writeUint32(w io.Writer, order binary.ByteOrder, i uint32) error {
	buf := make([]byte, 4)
	order.PutUint32(buf, i)
	return write(w, buf)
}

func writeUint64(w io.Writer, order binary.ByteOrder, i uint64) error {
	buf := make([]byte, 8)
	order.PutUint64(buf, i)
	return write(w, buf)
}

func writeString(w io.Writer, order binary.ByteOrder, s string) error {
	if err := writeUint16(w, order, uint16(len(s))); err != nil {
		return fmt.Errorf("write length: %w", err)
	}
	return write(w, []byte(s))
}

func writeFloat32(w io.Writer, order binary.ByteOrder, i float32) error {
	return writeUint32(w, order, math.Float32bits(i))
}

func writeFloat64(w io.Writer, order binary.ByteOrder, i float64) error {
	return writeUint64(w, order, math.Float64bits(i))
}
