package main

import (
	"io"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			io.Copy(os.Stdout, conn)
		}()
		defer conn.Close()
	}
}
