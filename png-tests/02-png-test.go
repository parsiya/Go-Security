// Using image/png to read PNG data and store it.

package main

import (
	"image/png"
	"os"
)

func main() {
	myPNG, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	defer myPNG.Close()

	img, err := png.Decode(myPNG)
	if err != nil {
		panic(err)
	}

	outfile, err := os.Open("out-png.txt")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	outfile.Write(img)
}
