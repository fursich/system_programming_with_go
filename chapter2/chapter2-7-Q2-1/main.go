package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hoge.txt")
	if err != nil {
		panic(err)
	}
	n, dept, height := 1, "総務部", 172.50
	fmt.Fprintf(file, "社員番号%d（所属: %s） 身長%f(cm)", n, dept, height)
}
