package main

import "fmt"

func Swap (p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Printf("*p1 = %d, *p2 = %d \n", *p1, *p2)
}

func main () {
	a, b := 10, 20
	Swap(&a, &b)
	fmt.Printf("a = %d, b = %d \n", a, b)
}
