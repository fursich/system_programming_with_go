package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		return
	}

	source := flag.Arg(0)
	dest := flag.Arg(1)
	// fmt.Printf("%v %v", org, org2)
	// flag.Args

	file, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	newFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer newFile.Close()

	io.Copy(newFile, file)
}
