package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	defer wg.Wait()
	wg.Add(1)
	fmt.Println("world")
	//go func() {
	//	defer wg.Done()
	//	for {
	//		fmt.Println("goroutine")
	//	}
	//}()
	go Hello()
}

func Hello() {
	defer wg.Done()
	fmt.Println("hello")
}
