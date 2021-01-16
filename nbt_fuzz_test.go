package nbt

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestFuzzCorpus(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "testdata/fuzz/corpus")
	if err := afero.Walk(fs, "", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		t.Run("fuzz="+path, func(t *testing.T) {
			f, err := fs.Open(path)
			if err != nil {
				panic(err)
			}
			defer func() { _ = f.Close() }()

			dec := NewDecoder(f, binary.BigEndian)
			tag, err := dec.ReadTag()
			if err != nil {
				if tag != nil {
					panic("tag != nil on error")
				}
				return // errors are ok, as long as the test doesn't panic
			}

			enc := NewEncoder(ioutil.Discard, binary.BigEndian)
			err = enc.WriteTag(tag)
			if err != nil {
				panic(err) // must be able to write tags that could be decoded
			}
		})
		return nil
	}); err != nil {
		t.Errorf("walk: %s", err)
	}
}
