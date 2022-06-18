package main

import "fmt"

type Person struct {
	name string
	age int
	sex byte
}

type Student struct {
	*Person
	id int64
	addr string
}

func main () {
	s1 := Student(&Person{"zeng", 19, 'm'} , 20202524, "cs")
	fmt.Println("s1 = ", s1)

	var s2 Student
	s2.Person = new(Person)
	s2.name = "wen"
	s2.age = 18
	s2.sex = 'f'
	s2.id = 2524
	s2.addr = "hy"
	fmt.Println(s2.name, s2.age, s2.sex, s2.id, s2.addr)
}

