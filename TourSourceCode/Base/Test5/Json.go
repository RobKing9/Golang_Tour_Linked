package main

import(
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string
	Subjects []string
	IsOK bool
	Price float64
}

func main () {
	s := IT{"itcast", []string{"Go", "C++", "Python"}, true, 444.444}
	buf, err := json.MarshalIndent(s, "", " ")
	if err != nil{
		fmt.Println("err= ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}