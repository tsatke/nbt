// +build gofuzz

package nbt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
)

func Fuzz(data []byte) int {
	dec := NewDecoder(bytes.NewReader(data), binary.BigEndian)
	tag, err := dec.ReadTag()
	if err != nil {
		if tag != nil {
			panic("tag != nil on error")
		}
		if errors.Is(err, io.EOF) {
			return 0
		}
		return 0
	}

	enc := NewEncoder(ioutil.Discard, binary.BigEndian)
	err = enc.WriteTag(tag)
	if err != nil {
		panic(err) // must be able to write tags that could be decoded
	}
	return 1
}
