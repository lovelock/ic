package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type RemoteHost struct {
	ShortName    string
	Host         string
	IdentityFile string
	User         string
	Port         string
}

func ListHosts() []string {
	m := ListBlocks()
	r := []string{}

	for k, _ := range m {
		r = append(r, k)
	}

	return r
}

func ListBlocks() map[string]RemoteHost {
	content := GetConfigContent()

	blocks := strings.Split(content, "\n\n")

	m := make(map[string]RemoteHost)

	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		var item RemoteHost
		for _, line := range lines {
			pair := strings.Split(line, " ")
			if len(pair) > 1 {
				k := pair[0]
				v := pair[1]
				if k == "Host" {
					item.ShortName = v
				}
				if k == "HostName" {
					item.Host = v
				}

				if k == "User" {
					item.User = v
				}

				if k == "Port" {
					item.Port = v
				}

				if k == "IdentityFile" {
					item.IdentityFile = v
				}
			}
		}

		if item.ShortName != "*" {
			m[item.ShortName] = item
		}
	}

	return m
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
