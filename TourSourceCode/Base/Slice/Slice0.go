package main

import (
	"fmt"
)

func main () {
	a := [5]int{1,3,4,7,0}
	s := a[1 : 4 : 5]
	fmt.Println("s = ",s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s1 := []int{1, 3, 5, 7}
	fmt.Println("s1 = ", s1)

	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
