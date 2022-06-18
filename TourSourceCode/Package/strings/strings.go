package main

import (
	"fmt"
	"strings"
)

var (
	s1 = "abaaa"
	s2 = "ccbbcc"
	s3 = []string{"hello", "world"}
)

func main() {
	//输出到字符串
	res1 := fmt.Sprintf("%s%s", s1, s2) //abaaaccbbcc
	fmt.Println(res1)
	//分离字符串
	res2 := strings.Split(s1, "b") //[]string [a aaa]
	fmt.Printf("%T, %v", res2, res2)
	fmt.Println()
	//是否包含子字符串
	b1 := strings.Contains(s1, "a") //true
	b2 := strings.Contains(s1, "d") //false
	fmt.Println(b1, b2)
	//连接字符串
	res3 := strings.Join(s3, "-") //hello-world
	fmt.Println(res3)
	//字符串的替换
	res4 := strings.Replace(s2, "c", "#", -1) //##bb##
	fmt.Println(res4)
	//字符串的下标
	res5 := strings.Index(s1, "b")     //1
	res6 := strings.LastIndex(s2, "b") //3
	fmt.Println(res5, res6)
	//去掉字符串s中首部以及尾部与字符串cutset中每个相匹配的字符
	res7 := strings.Trim("babacdb", "b") // abacd
	fmt.Println(res7)
}
