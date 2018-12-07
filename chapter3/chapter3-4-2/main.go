package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/koji/workspace/go/system programming/chapter3/chapter3-4-2/main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
