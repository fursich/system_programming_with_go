package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"net/http"
)

var fileName string

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	source := map[string]string{
		"hoge":   "huga",
		"Hello":  "World",
		"fooooo": "barrrrrr",
	}

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	writer, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}

	jsonEncoder := json.NewEncoder(writer)
	jsonEncoder.SetIndent("", "    ")
	jsonEncoder.Encode(source)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		return
	}

	fileName = flag.Arg(0)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
