// [urfave-cli](https://github.com/urfave/cli) also supports flags.

// godoc: [https://godoc.org/github.com/urfave/cli](https://godoc.org/github.com/urfave/cli)

package main

import (
	"fmt"
	"os"
	"sort"

	// go get before using it
	"github.com/urfave/cli"
)

var (
	// Variable to hold flags
	flags []cli.Flag

	// Variable to hold the host flag
	host string

	// Variable to hold the port flag
	port int
)

// Subcommands can be processed in init.
func init() {
	// Defining flags
	flags = []cli.Flag{
		// Creating a string flag which is the most common
		cli.StringFlag{
			// Alternate flags are separated by commas in Name.
			// We can call -t, --t, -target, --host, etc.
			Name: "t, target, host",
			// Default value for flag (optional) is in Value.
			Value: "127.0.0.1",
			// Default placeholder in -h usage string is set with ``:
			// "-t HOST, --target HOST, --host HOST
			//     hacking HOST (default: "127.0.0.1")"
			Usage: "hacking `HOST`",
			// host will hold the value of flag -host
			Destination: &host,
		},
		// Creating an int flag is similar
		cli.IntFlag{
			Name:        "p, port",
			Usage:       "target `PORT`",
			Destination: &port,
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

	// Hide version in usage
	app.HideVersion = true

	// Set flags
	app.Flags = flags

	// Called if no arguments are provided at runtime.
	// Expects type cli.ActionFunc == "func (*cli.Context) error"
	app.Action = noArgs

	// Sort flags for printing if needed, this will sort flags
	// before printing in usage
	sort.Sort(cli.FlagsByName(app.Flags))

	// Run the app
	app.Run(os.Args)
}

// flags can be processed in noArgs
func noArgs(c *cli.Context) error {
	// NumFlags() returns the number of flags set.
	if c.NumFlags() < 2 {
		// If not enough flags are set, print usage and exit
		cli.ShowAppHelp(c)
		return cli.NewExitError("please set both flags", 2)
	}
	// Similar to thec;s flag package, if flags are not set, they
	// will have a default value. They need to be checked
	// individually.
	fmt.Printf("hacking %s:%d", host, port)
	return nil
}
