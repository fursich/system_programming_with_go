package main

import (
	"fmt"
	"io"
	"os"
)

func open() {
	file, err := os.Create("foo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content2\n")
}

func read(s string) {
	file, err := os.Open("foo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(s)
	io.Copy(os.Stdout, file)
}

func append() {
	file, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Appended content\n")
}

func main() {
	open()
	read("Reading created file:")
	append()
	read("Reading appended file:")
}
