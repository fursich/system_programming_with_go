package main

import (
	"flag"
	"io"
	"os"
)

func CopyN(dest io.Writer, source io.Reader, length int64) {
	limitReader := io.LimitReader(source, length)
	io.Copy(dest, limitReader)
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		return
	}

	sourceFile := flag.Arg(0)
	targetFile := flag.Arg(1)

	source, err := os.Open(sourceFile)
	defer source.Close()
	if err != nil {
		panic(err)
	}

	target, err := os.Create(targetFile)
	defer target.Close()
	if err != nil {
		panic(err)
	}

	CopyN(target, source, 10)
	target.WriteString("\n\n<<<<HOGE hoge HOGE!!>>>>\n\n")
	CopyN(target, source, 10)
}
