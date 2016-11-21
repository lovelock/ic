package main

import (
	"fmt"
	"os"
	"os/user"
)

func SshConfigFile() string {
	user, err := user.Current()
	CheckError(err)
	return fmt.Sprintf("%s/.ssh/config", user.HomeDir)
}

func MyName() string {
	user, err := user.Current()
	CheckError(err)

	return user.Name
}

func DefaultPrivateKey() string {
	user, err := user.Current()
	CheckError(err)
	return fmt.Sprintf("%s/.ssh/id_rsa", user.HomeDir)
}

func RemoveRemoteHost(k string) map[string]RemoteHost {
	m := ListBlocks()
	delete(m, k)

	c := GetInitSetting()

	for _, r := range m {
		c = append(c, RecoverBlockLines(r)...)
	}

	Write(SshConfigFile(), c)

	nm := ListBlocks()
	return nm
}

func AddRemoteHost(host string, hostName string, privateKey string, port string, username string) RemoteHost {
	if hostName == "" {
		fmt.Println("Please porivde host at least")
		os.Exit(1)
	}

	if host == "" {
		host = hostName
	}

	if port == "" {
		port = "22"
	}

	if username == "" {
		username = MyName()
	}

	if privateKey == "" {
		privateKey = DefaultPrivateKey()
	}

	aRemoteHost := RemoteHost{
		Host:         host,
		HostName:     hostName,
		IdentityFile: privateKey,
		User:         username,
		Port:         port,
	}

	aSeg := fmt.Sprintf("\n\nHost %s\nHostName %s\nPort %s\nUser %s\nIdentityFile %s", aRemoteHost.Host, aRemoteHost.HostName, aRemoteHost.Port, aRemoteHost.User, aRemoteHost.IdentityFile)
	err := Append(SshConfigFile(), []byte(aSeg))
	CheckError(err)

	return aRemoteHost
}

func GetInitSetting() []byte {
	return []byte("Host *\nControlMaster auto\nControlPath ~/.ssh/master-%r@%h:%p\n")
}

func RecoverBlockLines(r RemoteHost) []byte {
	return []byte("\n\nHost " + r.Host + "\nHostName " + r.HostName + "\nPort " + r.Port + "\nUser " + r.User + "\nIdentityFile " + r.IdentityFile + "\n")
}

func WriteRemoteHostToBlock(r RemoteHost) []byte {
	lines := []byte("\n\nHost " + r.Host + "\nHostName " + r.HostName + "\nPort " + r.Port + "\nUser " + r.User + "\nIdentityFile " + r.IdentityFile + "\n")

	if r.Forwards != nil {
		for _, f := range r.Forwards {
			lines = append(lines, RecoverForwardLine(f)...)
		}
	}

	return lines
}
