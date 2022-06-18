package main

import "fmt"

type Student struct {
	id int64
	name string
	sex string
	age int
	addr string
}

func main () {
	//var s Student
	//var p *Student
	p := new(Student)
	//p = &s
	p.id = 20202524
	p.name = "zengxiangwen"
	p.sex = "male"
	p.age = 19
	p.addr = "cs"
	fmt.Println("p = ", p)
}
