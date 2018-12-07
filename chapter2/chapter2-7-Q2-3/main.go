package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
		"Fizz":  "Buzz",
	}

	gzipWriter := gzip.NewWriter(w)
	writer := io.MultiWriter(gzipWriter, os.Stdout)
	jsonEncoder := json.NewEncoder(writer)
	jsonEncoder.SetIndent("", "    ")
	jsonEncoder.Encode(source)

	gzipWriter.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
