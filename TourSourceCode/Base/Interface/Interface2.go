package main

import "fmt"

type Hunmaner interface {
	sayhi()
}

type Personer interface {
	Hunmaner
	sing(lrc string)
}
type Student struct{
	name string
	id int64
}

type Teacher struct {
	Student
	group string
}


func (temp *Student) sayhi() {
	fmt.Printf("Student[%s, %d\n] sayhello", temp.name, temp.id)
}

func (temp *Student) sing(lrc string) {
	fmt.Println("Student在唱着: ", lrc)
}

func (temp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %d, %s\n] sayhelloStudent", temp.name, temp.id, temp.group)
}

func main () {
	var h Hunmaner
	var hpro Personer
	hpro = &Student{"fumin", 20212703}
	h = hpro
	h.sayhi()
	hpro.sing("music")

	t := &Teacher{Student{"wen", 20202524}, "good"}
	h = t
	h.sayhi()
}
