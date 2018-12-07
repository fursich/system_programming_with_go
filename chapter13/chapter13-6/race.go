package main

import "fmt"

func main() {
	status := make(chan int)
	x := 0
	go func() {
		for i := 0; i < 10; i++ {
			x += 3
			fmt.Println(x)
		}
		status <- x
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(x)
		x++
	}
	fmt.Println(<-status)
}
