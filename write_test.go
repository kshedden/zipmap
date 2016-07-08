package zipmap

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	fname := "tmp.gob.gz"
	m1 := map[int64]float64{1: 35, 2: 1.5}
	Write(m1, fname)
	m2, err := ReadInt64Float64(fname)
	if err != nil {
		panic(err)
	}

	if fmt.Sprintf("%T", m1) != fmt.Sprintf("%T", m2) {
		t.Fail()
	}

	if len(m1) != len(m2) {
		t.Fail()
	}

	for k, _ := range m1 {
		if m1[k] != m2[k] {
			t.Fail()
		}
	}
}

func Test2(t *testing.T) {

	fname := "tmp.gob.gz"
	m1 := map[int64]int64{1: 35, 2: 5}
	Write(m1, fname)
	m2, err := ReadInt64Int64(fname)
	if err != nil {
		panic(err)
	}

	if fmt.Sprintf("%T", m1) != fmt.Sprintf("%T", m2) {
		t.Fail()
	}

	if len(m1) != len(m2) {
		t.Fail()
	}

	for k, _ := range m1 {
		if m1[k] != m2[k] {
			t.Fail()
		}
	}
}

func Test3(t *testing.T) {

	fname := "tmp.gob.gz"
	m1 := map[int64]string{1: "cat", 2: "dog"}
	Write(m1, fname)
	m2, err := ReadInt64String(fname)
	if err != nil {
		panic(err)
	}

	if fmt.Sprintf("%T", m1) != fmt.Sprintf("%T", m2) {
		t.Fail()
	}

	if len(m1) != len(m2) {
		t.Fail()
	}

	for k, _ := range m1 {
		if m1[k] != m2[k] {
			t.Fail()
		}
	}
}
