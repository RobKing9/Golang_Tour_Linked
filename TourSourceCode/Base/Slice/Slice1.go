package main

import "fmt"

func main () {
	s1 := []int{1, 3, 5}
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d \ncap = %d\n", len(s1), cap(s1))
	s1 = append (s1, 2)
	s1 = append(s1, 9)
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d \ncap = %d\n", len(s1), cap(s1))
}