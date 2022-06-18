package main

import "fmt"

func add(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main () {
	f := add (100)
	fmt.Println( f(50) )
	fmt.Println( f(100) )
}
