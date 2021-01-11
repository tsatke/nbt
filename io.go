package nbt

import (
	"fmt"
	"io"
)

func write(w io.Writer, buf []byte) error {
	return bufferedStreamOp("write", w.Write, buf)
}

func read(rd io.Reader, buf []byte) error {
	return bufferedStreamOp("read", rd.Read, buf)
}

func bufferedStreamOp(opName string, fn func([]byte) (int, error), buf []byte) error {
	n, err := fn(buf)
	if err := mustLen(len(buf), n, err); err != nil {
		return fmt.Errorf("%s: %w", opName, err)
	}
	return nil
}

func mustLen(len, n int, err error) error {
	if err != nil {
		return err
	}
	if len != n {
		return fmt.Errorf("expect %d bytes, but got %d", len, n)
	}
	return nil
}
