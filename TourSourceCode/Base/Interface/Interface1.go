package main

import "fmt"

//type People interface{
//	Speak(string) string
//}

type Student struct{}

func (s *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大傻逼"
	} else {
		talk = "hello"
	}
	return
}

func main () {
	//var peo People
	var s1 = &Student{}
	//peo = s1
	think := "sp"
	fmt.Println(s1.Speak(think))
}