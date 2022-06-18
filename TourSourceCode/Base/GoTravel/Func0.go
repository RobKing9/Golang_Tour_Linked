package main

import "fmt"

func Add(a, b int) int{
	return a+b
}
func Swap(x, y string) (string, string){
	return y, x
}

var java, python = true, false

func main () {
	fmt.Println(Add(123, 456))
	c, d := Swap("wen", "min")
	fmt.Println(c,d)
	fmt.Println(java, python)
}
