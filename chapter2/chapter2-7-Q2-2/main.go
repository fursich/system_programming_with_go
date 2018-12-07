package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {
	file, err := os.Create("hoge.csv")
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(file, os.Stdout)
	writer := csv.NewWriter(multiWriter)
	writer.Write([]string{"hoge 100 huga 200 funa foo bar baz"})
	writer.Flush()
}
