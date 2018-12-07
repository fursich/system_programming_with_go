package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	charA := io.NewSectionReader(programming, 5, 1)
	charS := io.LimitReader(system, 1)
	charC := io.LimitReader(computer, 1)
	charI := io.NewSectionReader(programming, 8, 1)
	charI2 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(charA, charS, charC, charI, charI2)
	io.Copy(os.Stdout, stream)
}
