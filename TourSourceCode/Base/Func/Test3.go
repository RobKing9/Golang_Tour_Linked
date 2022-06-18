package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	rand.Seed( time.Now().UnixNano())

	for i := 1; i<5000; i++{
		fmt.Println("rand = ", rand.Intn(100))
	}
	fmt.Println( time.Now().Format("2006-01-02 15:04:05"))
}
