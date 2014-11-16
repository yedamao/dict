package rw

import (
	"testing"
)

func TestWriteLine1(t *testing.T) {
	err := WriteLine("this is test", "test.dict")
	if err != nil {
		err.Error()
		t.Error()
	}
}

func TestWriteLine2(t *testing.T) {
	err := WriteLine("second line", "test.dict")
	if err != nil {
		err.Error()
		t.Error()
	}
}

func TestWriteLine3(t *testing.T) {
	err := WriteLine("this is third", "./test.dict")
	if err != nil {
		err.Error()
		t.Error()
	}
}

func TestWriteLine4(t *testing.T) {
	err := WriteLine("this is test", "../test.dict")
	if err != nil {
		err.Error()
		t.Error()
	}
}

func TestWrite(t *testing.T) {
	content := "func TestWriteLine4(t *testing.T) {}}"
	err := Write(content, "./code.txt")
	if err != nil {
		err.Error()
		t.Error()
	}
}
