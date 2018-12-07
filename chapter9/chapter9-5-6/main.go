package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	matched, err := filepath.Match("main*.go", "main-100.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(matched)

	files, err := filepath.Glob("../**/*.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
