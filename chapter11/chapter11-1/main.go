package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())
	fmt.Printf("parent process ID: %d\n", os.Getppid())

	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stdout, "Group ID: %d session ID: %d\n",
		syscall.Getpgrp(), sid)
}
