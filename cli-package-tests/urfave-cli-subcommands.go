// [urfave-cli](https://github.com/urfave/cli) is a popular CLI package.
// In this example we will show how to use it to create sub commands and
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
	// Variable to hold subcommands
	cmds []cli.Command
)

// Subcommands can be processed in init.
func init() {
	// Defining subcommands
	cmds = []cli.Command{
		{
			// Name of subcommand
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

	// Set name of the program
	app.Name = "AppName"

	// Set application usage
	app.Usage = "application usage"

	// Set application description
	app.Description = "description two"

	// Set application version - if not set version is "0.0.0"
	app.Version = "1.2.3.4"
	// Hide version in usage
	app.HideVersion = true

	// Set author name
	app.Author = "Author1"
	// Set author email
	app.Email = "author1@example.com"

	// For multiple authors we need to create a slice of cli.Author
	// Author has two fields, Name and Email (both string)
	// First Authors list will be printed and then Author
	app.Authors = []cli.Author{
		{"Author2", "author2@example.com"},
		{"Author3", "author3@example.com"},
	}

	// Set subcommands
	app.Commands = cmds

	// Called if no arguments are provided at runtime.
	// Expects type cli.ActionFunc == "func (*cli.Context) error"
	app.Action = noArgs

	// Run the app
	app.Run(os.Args)
}

// noArgs will run if no arguments are provided
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
		fmt.Printf("%sing %s\n", cmdName, t)
		return nil
	}
	// If there are no arguments, show help for that specific subcommands
	// and then return with an error.
	cli.ShowSubcommandHelp(c)
	return cli.NewExitError("no arguments for subcommand", 3)
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
	return cli.NewExitError("no arguments for subcommand", 3)
}
