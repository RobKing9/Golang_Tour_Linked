package main

import (
	"fmt"
	"os" //用来针对用户所用的操作系统的相关操作
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("当前路径:", dir) //当前路径: G:\java_workspace\test_os

	value := os.Getenv("path")
	fmt.Println(value) //获取环境变量的path值

}
