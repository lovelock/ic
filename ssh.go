package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func Ssh(username string, password string, ip string, port string) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	conn, err := ssh.Dial("tcp", ip+":"+port, config)

	if err != nil {
		log.Fatal("Failed to dail: ", err)
	}

	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}

	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	if err := session.Shell(); err != nil {
		log.Fatal("Failed to start a shell: ", err)
	} else {
		fmt.Println("Connect successfully")
	}

	err = session.Wait()
	if err != nil {
		log.Fatal("return")
	}
}
