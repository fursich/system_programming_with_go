package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		return
	}
	sec, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	fmt.Printf("now counting up %d second(s)..\n", sec)

	<-time.After(time.Duration(sec) * time.Second)

	fmt.Println("timer has expired\n")
}
