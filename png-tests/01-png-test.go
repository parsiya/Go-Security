// Read PNG image data using "image."
// "image/png" should be imported for image.Decode to recognize PNG files.

package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func main() {
	myPNG, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	defer myPNG.Close()

	_, imgType, err := image.Decode(myPNG)
	if err != nil {
		panic(err)
	}

	// fmt.Println(imgData)
	fmt.Println(imgType)

}
