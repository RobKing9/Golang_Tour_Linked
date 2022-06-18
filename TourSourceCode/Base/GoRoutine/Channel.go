package main

import (
	"fmt"
	"time"
)

func main () {

	ch := make (chan string)
	defer fmt.Println("主协程调佣完毕")
	go func () {
		for i := 0; i<3; i++ {
			defer fmt.Println("子协程调佣完毕")
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}

		ch <- "我是子协程"
	}()

	str := <-ch
	fmt.Println("str = ", str)
}