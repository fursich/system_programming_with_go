package main

import (
	"net"
	"net/http"
)

func writToConn(responses chan *http.Response, conn net.Conn) {
	defer conn.Close()
	for response := range responses {
		response.Write(conn)
	}
}

func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	for sessionResponse := range sessionResponses {
		response := <-sessionResponse
		response.Write(conn)
	}
}

func main() {
}
