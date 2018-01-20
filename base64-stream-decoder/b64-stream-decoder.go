// Example to demonstrate using base64 stream decoder with large files.
// Blog: https://parsiya.net/blog/2018-01-19-decoding-large-base64-files-with-go/
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// This creates a flag parameter. Meaning we can call -file or --file.
	var filename string
	flag.StringVar(&filename, "file", "", "input file")
	flag.Parse()

	fmt.Println("reading input file, this may take a while")

	// Open input file
	input, err := os.Open(filename)
	// We are panic-ing with every error because we want to stop if things
	// go wrong.
	if err != nil {
		panic(err)
	}
	// Close input file
	defer input.Close()

	// Open output file
	output, err := os.Create(filename + "-out")
	if err != nil {
		panic(err)
	}
	// Close output file
	defer output.Close()

	// Create base64 stream decoder from input file. *io.File implements the
	// io.Reader interface. In other words we can pass it to NewDecoder.
	decoder := base64.NewDecoder(base64.StdEncoding, input)
	// Magic! Copy from base64 decoder to output file
	io.Copy(output, decoder)

	fmt.Println("storing base64 decoder input file")
}
