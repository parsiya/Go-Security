// urfave-cli example #5
// Showing subcommand categories and using cli.Context passed to Action func.
// This time we are using the same function for Action of all sub commands
// and then grabbing the subcommand in the function.
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var cmds []cli.Command

func init() {
	cmds = []cli.Command{
		{
			// Name
			Name:    "hack",
			Aliases: []string{"haxorx", "pwn"},
			Usage:   "hack target",
			// Same function for all actions
			Action:   act,
			Category: "Hacking",
		},
		{
			Name:     "crack",
			Aliases:  []string{"c"},
			Usage:    "crack target",
			Action:   act,
			Category: "Hacking",
		},
		{
			Name:     "scan",
			Aliases:  []string{"s"},
			Usage:    "scan target",
			Action:   act,
			Category: "Recon",
		},
		{
			Name:     "probe",
			Aliases:  []string{"p"},
			Usage:    "probe target",
			Action:   act,
			Category: "Recon",
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Commands = cmds
	app.Action = noArgs
	app.Usage = "app to Hack the Planet!"
	app.HideVersion = true

	app.Run(os.Args)
}

// noArgs will run if no arguments/flags are provided
func noArgs(c *cli.Context) error {
	// Print app usage
	cli.ShowAppHelp(c)
	return nil
}

// act will detect the subcommand from c and print the appropriate message
// or subcommand help if applicable.
func act(c *cli.Context) error {
	if c.Args().Present() {
		t := c.Args().First()
		// c.Command.Name has command name
		// c.Command has a lot more info
		// https://godoc.org/github.com/dkolbly/cli#Command
		// Knowing the subcommand we can call different functions here.
		cmdName := c.Command.Name
		fmt.Printf("%sing %s", cmdName, t)
		return nil
	}
	cli.ShowSubcommandHelp(c)
	return nil
}

// No arguments
// $go run urfave-cli5.go
// NAME:
//    urfave-cli5.exe - app to Hack the Planet!
// USAGE:
//    urfave-cli5.exe [global options] command [command options] [arguments...]
// COMMANDS:
//      help, h  Shows a list of commands or help for one command
//    Hacking:
//      hack, haxorx, pwn  hack target
//      crack, c           crack target
//    Recon:
//      scan, s   scan target
//      probe, p  probe target
// GLOBAL OPTIONS:
//    --help, -h  show help

// No targets
// $ go run urfave-cli5.go hack
// NAME:
//    urfave-cli5.exe hack - hack target
// USAGE:
//    urfave-cli5.exe hack [arguments...]
// CATEGORY:
//    Hacking

// Hack the target
// $ go run urfave-cli5.go hack 127.0.0.1:1234
// hacking 127.0.0.1:1234
// $ go run urfave-cli5.go crack 127.0.0.1:1234
// cracking 127.0.0.1:1234
// Just adding -ing to the end does not work all the time
// $ go run urfave-cli5.go probe 127.0.0.1:1234
// probeing 127.0.0.1:1234
// Ditto
// $ go run urfave-cli5.go scan 127.0.0.1:1234
// scaning 127.0.0.1:1234
