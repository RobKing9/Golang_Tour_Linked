package main

import "fmt"

func main(){
	var s int
	fmt.Println("请输入一个正整数：")
	var n int
	fmt.Scanf("%d",n)
	for i:=0;i<=n;i++{
		s=s+f(i)
	}
	fmt.Printf("阶乘之和为:%d",s)
}

func f(n int) int{
	flag:=1
	for i:=0;i<=n;i++{
		flag*=i
	}
	return flag
}
