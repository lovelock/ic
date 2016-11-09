package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
	content := GetConfigContent()

	blocks := strings.Split(content, "\n\n")

	m := make(map[string]string)
	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			pair := strings.Split(line, " ")
			if len(pair) > 1 {
				for k, v := range pair {
					fmt.Println(k, ": ", v)
				}
			}
			m[k] = v
		}
	}

	//fmt.Println(m)
}

func GetConfigContent() string {
	configFile := SshConfigFile()
	fi, err := os.Open(configFile)
	CheckError(err)

	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	CheckError(err)

	return string(fd)
}
