package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 100
	//将int型转化为string
	istr := strconv.Itoa(i)
	fmt.Printf("%#v, %T\n", istr, istr)
	str := "1000"
	//将string转化为int
	strint, _ := strconv.Atoi(str)
	fmt.Printf("%#v, %T", strint, strint)

	// byte转数字
	s := "12345"           // s[0] 类型是byte
	num := int(s[0] - '0') // 1
	st := string(s[0])     // "1"
	b := byte(num + '0')   // '1'
	fmt.Printf("%d%s%c\n", num, st, b)
}
