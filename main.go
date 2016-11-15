package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
)

func main() {
	app := cli.NewApp()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Frost Wong",
			Email: "frostwong@gmail.com",
		},
	}
	app.Version = "1.0.0"
	app.Usage = "A ssh client for human"

	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Init config file for no request for password when one session is active",
			Action:  initAction,
		},
		{
			Name:         "connect",
			Aliases:      []string{"c"},
			Action:       connectRemoteHostAction,
			Usage:        "complete a task on the list",
			BashComplete: listAvailableHosts,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Action:  addRemoteHostAction,
			BashComplete: func(c *cli.Context) {
				options := []string{"--short_name", "--host", "--user_name", "--port", "--identity_file"}
				if c.NArg() > 0 {
					return
				}

				for _, o := range options {
					fmt.Println(o)
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "shortname,s",
					Usage: "A host's `SHORTNAME` that will be shown in the list",
				},
				cli.StringFlag{
					Name:  "username, u",
					Usage: "`User name` that will be used to log in remote host",
				},
				cli.StringFlag{
					Name:  "host",
					Usage: "Remote server's IP or `HOSTNAME`",
				},
				cli.StringFlag{
					Name:  "port, p",
					Usage: "The SSH `PORT NUMBER` that the remote host is listening to",
				},
			},
		},
		{
			Name:    "add_local_forward",
			Aliases: []string{"al"},
			Action: func(c *cli.Context) {
				short_name := c.String("short_name")
				localforward := c.String("local_forward")
				forward_ip := c.String("forward_ip")
				forward_port := c.String("forward_port")
				forward_user := c.String("forward_user")

				lPort := RandomPortForLocalForward()

				bs := ListBlocks()
				fmt.Println(bs)

				fmt.Println(short_name, forward_ip, forward_port)

				lSeg := fmt.Sprintf("\n\nHost %s\nHostName 127.0.0.1\nUser %s\nPort %d", localforward, forward_user, lPort)
				err := Append(SshConfigFile(), []byte(lSeg))

				CheckError(err)
			},
			BashComplete: func(c *cli.Context) {
				options := []string{"--short_name", "--identity_file", "--forward_ip", "--forward_port", "--forward_user"}
				if c.NArg() > 0 {
					return
				}

				for _, o := range options {
					fmt.Println(o)
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "local_forward, l",
					Usage: "LocalForward to SSH to another remote host",
				},
				cli.StringFlag{
					Name:  "identity_file,i",
					Usage: "`Private key` that is used to log in to remote host",
				},
				cli.StringFlag{
					Name:  "forward_ip",
					Usage: "IP that must be forwared",
				},
				cli.StringFlag{
					Name:  "forward_port",
					Usage: "Port that must be forwared",
				},
				cli.StringFlag{
					Name:  "forward_user",
					Usage: "Forward host's username",
				},
			},
		},
		{
			Name:         "remove",
			Aliases:      []string{"r"},
			Action:       removeRemoteHostAction,
			Usage:        "A host's `SHORTNAME` that will be shown in the list",
			BashComplete: listAvailableHosts,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Action:  listAvailableHosts,
			Usage:   "List all available remote hosts",
		},
	}

	app.BashComplete = func(c *cli.Context) {

		if c.NArg() > 0 {
			return
		}

		for _, i := range app.Commands {
			fmt.Println(i.Name)
		}
	}

	app.Run(os.Args)
}

func listAvailableHosts(c *cli.Context) {
	hosts := ListHosts()
	if c.NArg() > 0 {
		return
	}

	for _, h := range hosts {
		fmt.Println(h)
	}
}

func removeRemoteHostAction(c *cli.Context) {
	shortname := c.Args().Get(0)
	m := RemoveRemoteHost(shortname)
	_, ok := m[shortname]
	if ok {
		fmt.Printf("%s is removed from remote host list", shortname)
	} else {
		fmt.Printf("Remove %s from remote host list failed", shortname)
	}
}

func initAction(c *cli.Context) error {
	config := SshConfigFile()
	log.Printf("SSH config file is %s", config)

	if Exists(config) == false {
		err := Create(config)
		log.Printf("Create %s failed, verbose log %v", config, err)
		err = Append(config, GetInitSetting())
		CheckError(err)
	}

	return nil
}

func addRemoteHostAction(c *cli.Context) {
	shortname := c.String("s")
	host := c.String("host")
	port := c.String("p")
	username := c.String("u")
	private_key := c.String("i")

	if host == "" {
		fmt.Println("Please porivde host at least")
		os.Exit(1)
	}

	if shortname == "" {
		shortname = host
	}

	if port == "" {
		port = "22"
	}

	if username == "" {
		username = MyName()
	}

	if private_key == "" {
		private_key = DefaultPrivateKey()
	}

	aSeg := fmt.Sprintf("\n\nHost %s\nHostName %s\nPort %s\nUser %s\nIdentityFile %s", shortname, host, port, username, private_key)
	err := Append(SshConfigFile(), []byte(aSeg))
	CheckError(err)
}

func connectRemoteHostAction(c *cli.Context) error {
	host := c.Args().First()
	cmd := exec.Command("ssh", host)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Run()
	return nil
}
