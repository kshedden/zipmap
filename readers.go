package zipmap

import (
	"compress/gzip"
	"encoding/gob"
	"os"
)

func ReadInt64Float64(fname string) (map[int64]float64, error) {

	fid, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	rdr, err := gzip.NewReader(fid)
	if err != nil {
		return nil, err
	}
	dec := gob.NewDecoder(rdr)
	mp := make(map[int64]float64)
	err = dec.Decode(&mp)
	if err != nil {
		return nil, err
	}
	rdr.Close()
	fid.Close()

	return mp, nil
}

func ReadInt64Int64(fname string) (map[int64]int64, error) {

	fid, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	rdr, err := gzip.NewReader(fid)
	if err != nil {
		return nil, err
	}
	dec := gob.NewDecoder(rdr)
	mp := make(map[int64]int64)
	err = dec.Decode(&mp)
	if err != nil {
		return nil, err
	}
	rdr.Close()
	fid.Close()

	return mp, nil
}

func ReadInt64String(fname string) (map[int64]string, error) {

	fid, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	rdr, err := gzip.NewReader(fid)
	if err != nil {
		return nil, err
	}
	dec := gob.NewDecoder(rdr)
	mp := make(map[int64]string)
	err = dec.Decode(&mp)
	if err != nil {
		return nil, err
	}
	rdr.Close()
	fid.Close()

	return mp, nil
}
