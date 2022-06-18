package main

import "fmt"

type Dog struct{}

type Mover interface {
	move()
}

func (d Dog) move(){
	fmt.Println("狗会动")
}


func main (){
	var x Mover
	var baquan = Dog{} //baquan := Dog{}
	x = baquan
	var wangcai = &Dog{}
	x = wangcai
	x.move()
}
