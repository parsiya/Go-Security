// jessevdk/go-flags example #1 - INCOMPLETE
package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

// Create flags - based on the example in the package README.
var opts struct {
	// Bool slice is useful when creating things like -vvv
	Verbose []bool `short:"v" long:"verbose" description:"show verbose debug information"`
	Verbose []bool `short:"v2" long:"verbose2" description:"show verbose debug information"`

	// Host - using a string slice will allow us to pass multiple targets with
	// multiple -host(s). Not using a separate port flag this time.
	Host []string `short:"h" long:"host" description:"host to hack!"`

	// Random int - just to see what how it works.
	Timeout int `short:"t" long:"timeout" description:"timeout"`
}

func main() {
	fmt.Println(Verbose)
}
