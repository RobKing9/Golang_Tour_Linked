package main

import (
	"fmt"
	"time"
)

func modify( p *[5]int ) {
	(*p)[0] = 6
	fmt.Println(" modify *p = ",*p)
}

func main () {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	a := [5]int{1 : 4, 3 : 5, 4 : 20}
	modify(&a)
	fmt.Println("main : a = ",a)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}