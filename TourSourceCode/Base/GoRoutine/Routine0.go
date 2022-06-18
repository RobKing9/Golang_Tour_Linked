package main

import (
	"runtime"
	"fmt"
)

func main () {
	go func (){
			for i := 0;i < 5;i++{
				fmt.Println("hello go")
			}
	}()
	i := 0
	for {
		i++
		runtime.Gosched()
		fmt.Println("hello")
	}
}