package main

import "fmt"

func main () {
	var s1 Student = Student{20202524, "wen", "male", 19, "cs"}
	fmt.Println("s1 = ", s1)
}

type Student struct {
	id int64
	name string
	sex string
	age int
	addr string
}
