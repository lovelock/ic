package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
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
			Name:   "init",
			Usage:  "Init config file for no request for password when one session is active",
			Action: initAction,
		},
		{
			Name:         "connect",
			Action:       connectRemoteHostAction,
			Usage:        "complete a task on the list",
			BashComplete: listAvailableHosts,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Action:  addRemoteHostAction,
			BashComplete: func(c *cli.Context) {
				options := []string{"--Host", "--HostName", "--User", "--Port", "--IdentityFile"}

				if c.NArg() > 0 {
					return
				}

				for _, o := range options {
					fmt.Println(o)
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "Host",
					Usage: "A host's `name` that will be shown in the list",
				},
				cli.StringFlag{
					Name:  "User",
					Usage: "`User name` that will be used to log in remote host",
				},
				cli.StringFlag{
					Name:  "HostName",
					Usage: "Remote server's IP address or hostname",
				},
				cli.StringFlag{
					Name:  "Port",
					Usage: "The SSH `PORT NUMBER` that the remote host is listening to",
				},
				cli.StringFlag{
					Name:  "IdentityFile",
					Usage: "The private key that is used to connect to remote host",
				},
			},
		},
		{
			Name:   "info",
			Action: remoteHostInfo,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "Pretty",
					Usage: "Entry you want to describe",
				},
			},
			BashComplete: listAvailableHosts,
		},
		{
			Name: "add_local_forward",
			Action: func(c *cli.Context) {
				entry := c.String("Entry")
				host := c.String("Host")
				forwardIP := c.String("ForwardIP")
				forwardPort := c.String("ForwardPort")
				forwardUser := c.String("ForwardUser")
				identityFile := c.String("IdentityFile")

				localPort := RandomPortForLocalForward()

				thisEntry := RemoteHostDetail(entry)
				newForward := LocalForward{
					Host:       host,
					HostName:   forwardIP,
					LocalPort:  localPort,
					RemotePort: forwardPort,
					User:       forwardUser,
				}

				thisEntry.Forwards = append(thisEntry.Forwards, newForward)

				tmp := WriteRemoteHostToBlock(thisEntry)

				AddRemoteHost(host, "127.0.0.1", identityFile, localPort, forwardUser)
				RemoveRemoteHost(entry)
				Append(SshConfigFile(), tmp)
				ReconnectTunnel(entry)
			},
			BashComplete: func(c *cli.Context) {
				options := []string{"--Entry", "--Host", "--IdentityFile", "--ForwardIP", "--ForwardPort", "--ForwardUser"}
				if c.NArg() > 0 {
					return
				}

				for _, o := range options {
					fmt.Println(o)
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "Entry",
					Usage: "Which entry would be chosen to do local forwarding",
				},
				cli.StringFlag{
					Name:  "Host",
					Usage: "New entry's name",
				},
				cli.StringFlag{
					Name:  "IdentityFile",
					Usage: "`Private key` that is used to log in to remote host",
				},
				cli.StringFlag{
					Name:  "ForwardIP",
					Usage: "IP that must be forwared",
				},
				cli.StringFlag{
					Name:  "ForwardPort",
					Usage: "Port that must be forwared",
				},
				cli.StringFlag{
					Name:  "ForwardUser",
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

	if Exists(config) == false {
		err := Create(config)
		err = Append(config, GetInitSetting())
		CheckError(err)
	}

	return nil
}

func addRemoteHostAction(c *cli.Context) {
	host := c.String("Host")
	hostName := c.String("HostName")
	port := c.String("Port")
	username := c.String("User")
	privateKey := c.String("IdentityFile")

	aRemoteHost := AddRemoteHost(host, hostName, privateKey, port, username)
	fmt.Printf("%+v has been added\n", aRemoteHost)
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

func ReconnectTunnel(t string) {
	cmd := exec.Command("ssh", "-f", "-N", t)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func remoteHostInfo(c *cli.Context) {
	entry := c.Args().First()
	v := RemoteHostDetail(entry)
	fmt.Printf("Entry detail: %+v\n", v)
}
