package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ListHosts() []string {
	configFile := SshConfigFile()

	rw, err := os.Open(configFile)
	CheckError(err)

	defer rw.Close()

	rb := bufio.NewReader(rw)
	r := []string{}

	for {
		line, _, err := rb.ReadLine()
		if err == io.EOF {
			break
		}

		s := string(line)
		if strings.HasPrefix(s, "Host ") == true && strings.Contains(s, "*") == false {
			r = append(r, s[5:])
		}
	}

	return r
}

func ListBlocks() {
	configFile := SshConfigFile()
}
