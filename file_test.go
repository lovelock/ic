package main

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	name := "/tmp/mytestfile"
	err := Create(name)
	if err != nil {
		t.Errorf("Create file %s failed", name)
	}
}

func TestAppendFile(t *testing.T) {
	d1 := []byte("hello\nworld\n")
	Append("/tmp/Hello.java", d1)
}

func TestExists(t *testing.T) {
	r := Exists("/bin/bash")
	if r == false {
		t.Error("expected the file exists")
	}

	r = Exists("/bin/luanqibazao")
	if r == true {
		t.Error("expected the file does not exists")
	}
}
