package main

import (
	"testing"
)

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
