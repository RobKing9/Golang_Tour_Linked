package main

import "fmt"

type Student struct {
	name string
	age int
}

func (s *Student) Notify () {
	fmt.Printf("%v : %v\n", s.name, s.age)
}

func main () {
	s1 := Student{"fumin", 17}
	s1.Notify()

	s2 := Student{"zxw", 19}
	s3 := &s2
	s3.Notify()
}