package main

import (
	"fmt"
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

func GetInitSetting() []byte {
	return []byte("Host *\nControlMaster auto\nControlPath ~/.ssh/master-%r@%h:%p\n")
}

func RecoverBlockLines(r RemoteHost) []byte {
	return []byte("\n\nHost " + r.ShortName + "\nHostName " + r.Host + "\nPort " + r.Port + "\nUser " + r.User + "\nIdentityFile " + r.IdentityFile + "\n")
}
