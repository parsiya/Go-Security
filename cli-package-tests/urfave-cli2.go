// urfave-cli example #2
// Accessing non-flag arguments with c.Args()
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	// This function will be executed if no arguments are provided.
	// Should be of cli.ActionFunc type: func (*cli.Context) error
	app.Action = noArgs

	app.Run(os.Args)
}

// Print all arguments
func noArgs(c *cli.Context) error {
	// Returns true if there are any arguments present
	fmt.Println(c.Args().Present())

	// Get the first one
	fmt.Println(c.Args().First())

	// Read and print all non-flag arguments
	// c.NArg() returns the number of args == len(c.Args())
	if c.Args().Present() {
		for i := 0; i < c.NArg(); i++ {
			// Alternate: c.Args().Get(i) == c.Args()[i]
			fmt.Printf("Argument #%d: %s\n", i, c.Args().Get(i))
		}
	}
	return nil
}

// $ go run urfave-cli2.go 111 222 3333
// true
// 111
// Argument #0: 111
// Argument #1: 222
// Argument #2: 3333
