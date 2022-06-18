package main

import "fmt"

type User struct {
	name string
	id int64
}

func (self *User) Test () {
	fmt.Printf("%p, %v\n", self, self)
}

func main () {
	u := User{"fumin", 20212703}
	u.Test()

	mValue := u.Test
	mValue()

	mExpression := (*User).Test
	mExpression(&u)
}