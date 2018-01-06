// [urfave-cli](https://github.com/urfave/cli) is a popular CLI package.
// In this example we will show how to use it to create flags, sub commands and
// act on them.
// godoc: [https://godoc.org/github.com/urfave/cli](https://godoc.org/github.com/urfave/cli)
package main

import (
	"fmt"
	"os"

	// go get before using it
	"github.com/urfave/cli"
)

var (
	// Hold subcommands
	cmds []cli.Command
	// Hold flags
	flags []cli.Flag

	// Hold host
	host string
	// Hold port
	port int
)

func init() {
	// We will set flags, subcommands in init.

	// Set flags
	flags = []cli.Flag{
		cli.StringFlag{
			// Alternate flags are separated by commas.
			// We can call -t, --t, -target, --host, etc.
			Name: "t, target, host",
			// Default value for flag (optional)
			Value: "127.0.0.1",
			// Default placeholder in -h usage string is set with ``:
			// "-t HOST, --target HOST, --host HOST  hacking HOST (default: "127.0.0.1")"
			Usage: "hacking `HOST`",
			// Variable to hold the value of flag
			Destination: &host,
		},
		// IntFlag
		cli.IntFlag{
			Name:        "p, port",
			Value:       22,
			Usage:       "target `PORT`",
			Destination: &port,
		},
	}

	// Subcommands
	cmds = []cli.Command{
		{
			// Name
			Name: "hack",
			// Aliases (similar to alternate flags) stored in a string slice
			Aliases: []string{"haxorx", "pwn"},
			// Usage text - using `` to create placeholders is not supported here
			Usage: "hack target",
			// Function to call when this sub command is activated
			// Similar to app.Action, this function should be of type
			// cli.ActionFunc == func (*cli.Context) error
			Action: act,
			// Subcommand category
			Category: "Hacking",
		},
		{
			Name:    "crack",
			Aliases: []string{"c"},
			Usage:   "crack target",
			// It's possible to use the same function for different subcommands
			// (they do not have to be in the same category) and then detect
			// the command inside.
			Action:   act,
			Category: "Hacking",
		},
		{
			Name:    "scan",
			Aliases: []string{"s"},
			Usage:   "scan target",
			// Different function for this subcommand.
			Action:   scan,
			Category: "Recon",
		},
	}
}

func main() {
	// Create a new app
	app := cli.NewApp()

	// Name of the program
	app.Name = "AppName"

	// Application usage
	app.Usage = "application usage"

	// Description
	app.Description = "description two"

	// Version - if not version "0.0.0" will be printed
	app.Version = "1.2.3.4"
	// We can hide version instead
	app.HideVersion = true

	// Author name
	app.Author = "Author1"
	// Author email
	app.Email = "author1@example.com"

	// For multiple authors we need to create a slice of cli.Author
	// Author has two fields, Name and Email (both string)
	// First Authors list will be printed and then Author
	// AUTHORS:
	//    Author2 <author2@example.com>
	//    Author3 <author3@example.com>
	//    Author1 <author1@example.com>
	app.Authors = []cli.Author{
		{"Author2", "author2@example.com"},
		{"Author3", "author3@example.com"},
	}

	// Set subcommands
	app.Commands = cmds
	// Set flags
	app.Flags = flags

	// This function will be called if no arguments are provided at runtime
	// Expects type cli.ActionFunc == "func (*cli.Context) error"
	app.Action = noArgs

	// Run the app
	app.Run(os.Args)
}

// noArgs will run if no arguments/flags are provided
func noArgs(c *cli.Context) error {
	// Print app usage
	cli.ShowAppHelp(c)

	// It's possible to change the return status code and error here
	// cli.NewExitError creates a a new error and the return status code for
	// the application.
	return cli.NewExitError("no commands provided", 2)
}

// act will detect the subcommand from c and print the appropriate message
// or subcommand help if applicable.
func act(c *cli.Context) error {
	// Check for arguments after the subcommand
	if c.Args().Present() {
		// Get the next argument.Can also do c.Args().Get(0)
		// c.NArgs() returns the number of arguments and can be used in a for
		// along with .Get(i)
		t := c.Args().First()
		// c.Command.Name has command name and c.Command has a lot more [info](https://godoc.org/github.com/dkolbly/cli#Command)
		// Knowing the subcommand we can call different functions here.
		cmdName := c.Command.Name
		fmt.Printf("%sing %s", cmdName, t)
		return nil
	}
	// If there are no arguments, show help for that specific subcommands
	// and then return with an error.
	cli.ShowSubcommandHelp(c)
	return cli.NewExitError("no arguments provided for subcommand", 3)
}

// scan will "scan" the target
func scan(c *cli.Context) error {
	// Check for arguments after the subcommand
	if c.Args().Present() {
		// Get next argument (the target)
		t := c.Args().First()
		fmt.Println("scanning", t)
		return nil
	}
	// Show scan subcommand help
	cli.ShowSubcommandHelp(c)
	return cli.NewExitError("no arguments provided for subcommand", 3)
}
