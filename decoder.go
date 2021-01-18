package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// Decoder describes a component that is capable of
// reading NBT tags from some source.
type Decoder interface {
	ReadTag() (Tag, error)
}

type decoder struct {
	rd io.Reader
	bo binary.ByteOrder
}

// NewDecoder creates a new Decoder that will decode from the given reader and respect
// the given byte order. The byte order has to be compliant with the byte order of the
// NBT data in the source.
func NewDecoder(source io.Reader, byteOrder binary.ByteOrder) Decoder {
	return &decoder{
		rd: source,
		bo: byteOrder,
	}
}

func (d decoder) ReadTag() (Tag, error) {
	idByte, err := readByte(d.rd, d.bo)
	if err != nil {
		return nil, fmt.Errorf("read ID: %w", err)
	}
	id := ID(idByte)

	tag, err := newTagFromID(id)
	if err != nil {
		return nil, fmt.Errorf("new tag: %w", err)
	}
	if tag.ID() == IDTagEnd {
		return tag, nil
	}

	name, err := readString(d.rd, d.bo)
	if err != nil {
		return nil, fmt.Errorf("read tag name: %w", err)
	}
	tag.SetName(name)

	if err := tag.ReadFrom(d.rd, d.bo); err != nil {
		return nil, fmt.Errorf("read %s with name '%s' from: %w", tag.ID(), tag.Name(), err)
	}
	return tag, nil
}

func readByte(rd io.Reader, _ binary.ByteOrder) (byte, error) {
	buf := make([]byte, 1)
	if err := read(rd, buf); err != nil {
		return 0, err
	}
	return buf[0], nil
}

func readUint16(rd io.Reader, bo binary.ByteOrder) (uint16, error) {
	buf := make([]byte, 2)
	if err := read(rd, buf); err != nil {
		return 0, err
	}
	return bo.Uint16(buf), nil
}

func readUint32(rd io.Reader, bo binary.ByteOrder) (uint32, error) {
	buf := make([]byte, 4)
	if err := read(rd, buf); err != nil {
		return 0, err
	}
	return bo.Uint32(buf), nil
}

func readUint64(rd io.Reader, bo binary.ByteOrder) (uint64, error) {
	buf := make([]byte, 8)
	if err := read(rd, buf); err != nil {
		return 0, err
	}
	return bo.Uint64(buf), nil
}

func readString(rd io.Reader, bo binary.ByteOrder) (string, error) {
	strLen, err := readUint16(rd, bo)
	if err != nil {
		return "", fmt.Errorf("read length: %w", err)
	}

	buf := make([]byte, strLen)
	if err := read(rd, buf); err != nil {
		return "", err
	}
	return string(buf), nil
}

func readFloat32(rd io.Reader, bo binary.ByteOrder) (float32, error) {
	i, err := readUint32(rd, bo)
	return math.Float32frombits(i), err
}

func readFloat64(rd io.Reader, bo binary.ByteOrder) (float64, error) {
	i, err := readUint64(rd, bo)
	return math.Float64frombits(i), err
}

func read(rd io.Reader, buf []byte) error {
	_, err := io.ReadAtLeast(rd, buf, len(buf))
	return err
}
