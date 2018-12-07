package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `１行目
２行め
３行め`

func main() {
	// reader := bufio.NewReader(strings.NewReader(source))
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	fmt.Printf("%#v\n", line)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
