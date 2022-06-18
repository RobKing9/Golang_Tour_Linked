package main

import "fmt"

func test2(i int) int {
	if i==1{
		return 1
	}
	return i+test2(i-1)
}

func main(){
	var sum int
	sum=test2(100)
	fmt.Println("sum=",sum)
}
