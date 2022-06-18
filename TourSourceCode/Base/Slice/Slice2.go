package main

import (
	"math/rand"
	"fmt"
	"time"
)

func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0 ; i<len(s) ; i++ {
		s[i] = rand.Intn(100)
	}
}

func main () {
	//a := []int{1, 3, 4}
	//b := []int{8, 9}
	//copy(b, a)
	//fmt.Println("b = ",b)
	n := 10
	s := make([]int, n)
	InitData(s)
	fmt.Println("s = ", s)
}
