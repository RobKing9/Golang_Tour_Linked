package main

import (
	"fmt"
	"clac"
)
func init () {
	fmt.Println("this is main init")
}
func main () {
	a := calc.Add(10.23)
	fmt.Println("a = ",a)
}
