// zlib inflate (decompress).

package main

import (
	"compress/zlib"
	"io"
	"os"
)

func main() {
	zlibFile, err := os.Open("test.zlib")
	if err != nil {
		panic(err)
	}
	defer zlibFile.Close()

	r, err := zlib.NewReader(zlibFile)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	outFile, err := os.Create("out-zlib")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	io.Copy(outFile, r)
}
