package main

import (
	"fmt"
	"testing"
)

func TestSshConfigFile(t *testing.T) {
	f := SshConfigFile()
	fmt.Println(f)
}

func TestMyName(t *testing.T) {
	n := MyName()

	if n != "frost" {
		t.Error("My name is wrong")
	}
}
