package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	var a [10]int
	fmt.Println( len(a) )
	for i := 0; i < len(a); i++ {
		a[i] = i+1
		fmt.Printf("i = %d \na[i] = %d \n", i, a[i])
	}

	b := [5]int{1, 4, 6}
	fmt.Println("b = ", b)

	c := [5]int{0: 10,4: 50}
	fmt.Println("c = ", c)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
