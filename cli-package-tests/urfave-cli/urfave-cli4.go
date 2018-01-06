// urfave-cli example #4
// Using sub commands to hack the planet!
// Usually when using subcommands we pass/read another non-flag argument
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
			Name: "hack",
			// Aliases (similar to alternate flags) stored in a string slice
			Aliases: []string{"haxorx", "pwn"},
			// Usage text - using `` to create placeholders is not supported here
			Usage: "hack target",
			// Function to call when this sub command is activated
			// Similar to app.Action, this function should be of type
			// cli.ActionFunc == func (*cli.Context) error
			Action: hack,
		},
		{
			Name:    "scan",
			Aliases: []string{"s"},
			Usage:   "scan target",
			Action:  scan,
		},
	}
}

func main() {
	app := cli.NewApp()
	// Set commands
	app.Commands = cmds

	// This function will be executed if no arguments are provided.
	// Should be of cli.ActionFunc type: func (*cli.Context) error
	app.Action = noArgs
	app.Usage = "app to Hack the Planet!"

	// Hide version
	app.HideVersion = true

	app.Run(os.Args)
}

// If nothing is provided, print an error message
func noArgs(c *cli.Context) error {
	// Print app usage
	cli.ShowAppHelp(c)
	return nil
}

// hack will "hack" the target
func hack(c *cli.Context) error {
	// If there are arguments present
	if c.Args().Present() {
		// Get next argument (the target)
		t := c.Args().First()
		fmt.Println("hacking", t)
		return nil
	}
	// We can print subcommand help for this specific command
	cli.ShowSubcommandHelp(c)
	return nil
}

// scan will "scan" the target
func scan(c *cli.Context) error {
	// If there are arguments present
	if c.Args().Present() {
		// Get next argument (the target)
		t := c.Args().First()
		fmt.Println("scanning", t)
		return nil
	}
	cli.ShowSubcommandHelp(c)
	return nil
}
