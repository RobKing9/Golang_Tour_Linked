package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println(strings.Contains("helloworld", "hello"))

	s := []string{"abc", "hello", "mike", "go"}

	fmt.Println(strings.Join(s,"*"))   //Join

	fmt.Println(strings.Index("avc hello", "helo"))  //Indx

	fmt.Println(strings.Repeat("hello", 5))    //Repeat

}
