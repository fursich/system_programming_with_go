package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	storage := make([]byte, 25)
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
	buffer.WriteString("Buffer example2\n")

	size, err := buffer.Read(storage)
	if err != nil {
		panic(err)
	}
	fmt.Printf("size %d, string %s\n\n", size, storage)

	buffer.WriteString("Buffer example3\n")
	fmt.Println(buffer.String())
}
