// urfave-cli example #1
// See how different App properties show up in -h
package main

import (
	// "errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Create a new App
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

	// For multiple authors we need to create a slice of cli.Author.
	// Author has two fields, Name and Email (both string)
	app.Authors = []cli.Author{
		{"Author2", "author2@example.com"},
		{"Author3", "author3@example.com"},
	}

	// First Authors list will be printed and then Author
	// AUTHORS:
	//    Author2 <author2@example.com>
	//    Author3 <author3@example.com>
	//    Author1 <author1@example.com>

	// This function will be executed if no arguments are provided.
	// Expects type cli.ActionFunc which is "func (*cli.Context) error"
	app.Action = noSubs

	app.Run(os.Args)
}

// Print something and return
func noSubs(c *cli.Context) error {
	fmt.Println("inside app.Action")
	return nil

	// Seems like returning with a non-nil value has no effect - at least here
	// return errors.New("test error message")
}

// Output without parameters
// $ go run urfave-cli1.go
// inside app.Action

// Output for -h
// $ go run urfave-cli1.go -h
// NAME:
//    AppName - application usage

// USAGE:
//    urfave-cli1.exe [global options] command [command options] [arguments...]

// DESCRIPTION:
//    description two

// AUTHORS:
//    Author2 <author2@example.com>
//    Author3 <author3@example.com>
//    Author1 <author1@example.com>

// COMMANDS:
//      help, h  Shows a list of commands or help for one command

// GLOBAL OPTIONS:
//    --help, -h  show help
