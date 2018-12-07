package main

import (
	"crypto/rand"
	"flag"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		return
	}

	dest := flag.Arg(0)

	newFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.CopyN(newFile, rand.Reader, 1024)
}
