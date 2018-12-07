package main

import (
	"archive/zip"
	"flag"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		return
	}

	dest := flag.Arg(0)

	newFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	zipWriter := zip.NewWriter(newFile)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("hoge.txt")
	if err != nil {
		panic(err)
	}
	str := []byte("hogehogehoge hgoegheohgoe")
	writer.Write(str)
}
