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
			Action: func(c *cli.Context) error {
				config := SshConfigFile()
				log.Printf("SSH config file is %s", config)

				if Exists(config) == false {
					err := Create(config)
					log.Printf("Create %s failed, verbose log %v", config, err)
					err = Append(config, GetInitSetting())
					CheckError(err)
				}

				return nil
			},
		},
		{
			Name:    "connect",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) error {
				host := c.Args().First()
				cmd := exec.Command("ssh", host)
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr

				cmd.Run()
				return nil
			},
			Usage:        "complete a task on the list",
			BashComplete: listAvailableHosts,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Action: func(c *cli.Context) {
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
			},
			BashComplete: func(c *cli.Context) {
				// 列表要过滤掉已经添加的选项
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
				cli.StringFlag{
					Name:  "localforward, l",
					Usage: "LocalForward to SSH to another remote host",
				},
				cli.StringFlag{
					Name:  "IdentityFile,i",
					Usage: "`Private key` that is used to log in to remote host",
				},
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Action: func(c *cli.Context) {
				shortname := c.Args().Get(0)
				m := RemoveRemoteHost(shortname)
				_, ok := m[shortname]
				if ok {
					fmt.Printf("%s is removed from remote host list", shortname)
				} else {
					fmt.Printf("Remove %s from remote host list failed", shortname)
				}
			},
			Usage:        "A host's `SHORTNAME` that will be shown in the list",
			BashComplete: listAvailableHosts,
		},
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
