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

func TestRemoveRemoteHost(t *testing.T) {
	m := RemoveRemoteHost("fedora")
	if _, ok := m["fedora"]; ok {
		t.Error("Remove %s from remote host list failed", "fedora")
	}
}

func TestRecoverBlockLines(t *testing.T) {
	l := ListBlocks()
	fmt.Println(string(RecoverBlockLines(l["ubuntu"])))
}
