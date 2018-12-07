package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := &sync.Map{}

	smap.Store("hello", "world")
	smap.Store("test", "hoge")
	smap.Store(1, 2)

	smap.Delete("test")

	value, ok := smap.Load("test")
	fmt.Printf("key=%v, value=%v, exists?=%v\n", "test", value, ok)
	value, ok = smap.Load("hello")
	fmt.Printf("key=%v, value=%v, exists?=%v\n", "hello", value, ok)

	smap.LoadOrStore(1, 3)
	smap.LoadOrStore("test", "hogehoge")
	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

}
