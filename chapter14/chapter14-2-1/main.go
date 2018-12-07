package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputs := make(chan []byte)

	go func() {
		a, _ := ioutil.ReadFile("a.txt")
		inputs <- a
	}()

	go func() {
		b, _ := ioutil.ReadFile("b.txt")
		inputs <- b
	}()

	// (a heavy routine)

	fmt.Printf("%s\n", <-inputs)
	fmt.Printf("%s\n", <-inputs)
}
