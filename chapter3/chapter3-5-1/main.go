package main

import (
	"Strings"
	"io"
	"os"
)

func main() {
	reader := strings.NewReader("Example of io.SectioinReaer\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
