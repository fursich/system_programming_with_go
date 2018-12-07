package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: google.com\r\n\r\n")
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	fmt.Println(res.Header)
}
