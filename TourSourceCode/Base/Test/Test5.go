package main

import "fmt"

func main () {
	sum := func(a, b int) int {
		return a+b
	} (32,45)

	fmt.Printf("sum = %d\n", sum)

	test_fun := func(a, b int) int {
		return a+b
	}

	x := test_fun (32, 56)
	y := test_fun (43, 54)
	fmt.Printf("x = %d \ny = %d \n", x, y)
}