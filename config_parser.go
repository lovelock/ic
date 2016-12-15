package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type RemoteHost struct {
	Host         string
	HostName     string
	IdentityFile string
	User         string
	Port         string
	Forwards     []LocalForward
}

type LocalForward struct {
	User       string
	Host       string
	HostName   string
	LocalPort  string
	RemotePort string
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
			pair := strings.SplitN(line, " ", 2)
			if len(pair) > 1 {
				k := pair[0]
				v := pair[1]
				if k == "Host" {
					item.Host = v
				}
				if k == "HostName" {
					item.HostName = v
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

				if k == "LocalForward" {
					localForward := ParseLocalForward(v)
					item.Forwards = append(item.Forwards, localForward)
				}
			}
		}

		if item.Host != "*" {
			m[item.Host] = item
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

func RemoteHostDetail(entry string) RemoteHost {
	bs := ListBlocks()

	return bs[entry]
}

func ParseLocalForward(forward string) LocalForward {
	v1 := strings.Split(forward, " ")
	localPort := v1[0]
	hp := v1[1]
	v2 := strings.Split(hp, ":")
	remoteHost := v2[0]
	remotePort := v2[1]

	localForward := LocalForward{
		HostName:   remoteHost,
		LocalPort:  localPort,
		RemotePort: remotePort,
		User:       "wqc",
	}

	return localForward
}

func RecoverForwardLine(lf LocalForward) []byte {
	return []byte(fmt.Sprintf("LocalForward %s %s:%s\n", lf.LocalPort, lf.HostName, lf.RemotePort))
}
