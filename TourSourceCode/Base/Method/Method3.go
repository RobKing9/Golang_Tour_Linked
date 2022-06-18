package main

import "fmt"

type Student struct{
	name string
	age int
	sex byte
}

func (s Student) Value(a string, b int, c byte ){
	s.name = a
	s.age = b
	s.sex = c
}

func (s *Student) Pointer(a string, b int, c byte ){
	s.name = a
	s.age = b
	s.sex = c
}

func main () {
	s1 := Student{"wen", 19, 'm'}
	s1.Value("min", 17, 'f')
	fmt.Println(s1)
	s1.Pointer("min", 17, 'f')
	fmt.Println(s1)
}