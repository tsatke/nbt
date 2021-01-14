package nbt

import (
	"io"
)

func write(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	return err
}

func read(rd io.Reader, buf []byte) error {
	_, err := io.ReadAtLeast(rd, buf, len(buf))
	return err
}
