package main

import (
	//"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
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

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)

	if err != nil {
		panic(err)
	}

	defer terminal.Restore(fd, oldState)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	//in, _ := session.StdinPipe()

	termWith, termHeight, err := terminal.GetSize(fd)
	fmt.Println(termWith, termHeight)

	if err != nil {
		panic(err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm-256color", termWith, termHeight, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	if err := session.Shell(); err != nil {
		panic(err)
	}

	/*
	 *for {
	 *    reader := bufio.NewReader(os.Stdin)
	 *    str, _ := reader.ReadString('\n')
	 *    //fmt.Fprint(in, str)
	 *    session.Run(str)
	 *}
	 */
}
