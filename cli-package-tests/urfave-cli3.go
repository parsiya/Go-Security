// urfave-cli example #3
// Sets flags host and port then prints host:port.
package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

var (
	// host flag is a string
	host string

	// port is an int
	port int

	// Variable to hold all flags
	flags []cli.Flag
)

func init() {
	flags = []cli.Flag{
		cli.StringFlag{
			// Alternate flags are separated by commas
			Name: "t, target, host",
			// Default value (optional)
			Value: "127.0.0.1",
			// Default placeholder in -h usage string is set with ``:
			// "-t HOST, --target HOST, --host HOST  hacking HOST (default: "127.0.0.1")"
			Usage:       "hacking `HOST`",
			Destination: &host,
		},
		cli.IntFlag{
			Name:        "p, port",
			Value:       22,
			Usage:       "target `PORT`",
			Destination: &port,
		},
	}
}

func main() {
	app := cli.NewApp()
	// This function will be executed if no arguments are provided.
	// Should be of cli.ActionFunc type: func (*cli.Context) error
	app.Action = noArgs
	app.Flags = flags
	app.Usage = "app to Hack the Planet!"

	// Hide version
	app.HideVersion = true

	// Sort flags for printing if needed
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}

func noArgs(c *cli.Context) error {

	fmt.Printf("%s:%d", host, port)
	return nil
}
