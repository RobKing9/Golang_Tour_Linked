package main

import "fmt"

func main () {
	var a int = 10
	fmt.Printf("a = %d \n",a)
	fmt.Printf("&a = %v \n",&a)
	var p *int
	p = new (int)
	*p=334
	fmt.Printf("p = %v \n&a = %v \n*p = %d \n", p, &a, *p)
}
