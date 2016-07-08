package zipmap

import (
	"compress/gzip"
	"encoding/gob"
	"os"
)

// Write writes an arbitrary object to a compressed gob file with the
// given file name.
func Write(obj interface{}, fname string) error {

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fid.Close()
	wtr := gzip.NewWriter(fid)
	defer wtr.Close()
	enc := gob.NewEncoder(wtr)
	err = enc.Encode(obj)
	if err != nil {
		return err
	}

	return nil
}
