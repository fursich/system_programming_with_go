package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	// io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	// conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))

	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)

	io.Copy(os.Stdout, conn)
}
