package main

import "fmt"

func main(){
	var k int
	n:=0
	fmt.Println("请输入一个正整数： ")
	fmt.Scanf("%d",&k)
	for sum:=0;sum>k;n++{
		sum=sum+1/n
	}
	fmt.Printf("n最小为%d",n)
}