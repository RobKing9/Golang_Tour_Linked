package main

import (
	"fmt"
	"mypkg"
)
func init () {
	fmt.Println("this is main init")
}

func main () {
	mypkg.Myfunc()
}