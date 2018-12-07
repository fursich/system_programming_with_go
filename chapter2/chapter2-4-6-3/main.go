package main

import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-Test", "ヘッダーも追加できます")
	request.Write(os.Stdout)
	// os.Stdout.Write(request) // * error!
}
