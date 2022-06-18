package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	//count := len(os.Args)
	//fmt.Println("参数总个数:", count)
	//
	//fmt.Println("参数详情:")
	//for i := 0; i < count; i++ {
	//	fmt.Println(i, ":", os.Args[i])
	//}

	//name := flag.String("name", "zxw", "请输入名字")
	//age := flag.Int("age", 19, "年龄")
	//married := flag.Bool("married", false, "是否结婚")
	//flag.Parse()
	//fmt.Println(*name)
	//fmt.Println(*age)
	//fmt.Println(*married)
	//fmt.Println(flag.Args())
	//fmt.Println(flag.NFlag())
	//fmt.Println(flag.NArg())
}
