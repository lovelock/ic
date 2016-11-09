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
