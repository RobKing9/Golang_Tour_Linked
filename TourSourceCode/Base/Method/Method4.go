package main

import "fmt"

type Person struct{
	name string
	age int
	sex byte
}

type Student struct {
	Person
	id int64
}

func (temp *Person) print() {
	fmt.Printf("name = %s, age = %d, sex = %c\n", temp.name, temp.age, temp.sex)
}

func (temp *Student) print() {
	fmt.Printf("name = %s, age = %d, sex = %c, id = %d\n", temp.name, temp.age, temp.sex, temp.id)
}

func main () {
	s := Student{Person{"fumin", 17, 'f'}, 20212703}
	s.print()
	s.Person.print()
}