package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("hello GMP %d", i)
	}
	/*
		//创建trace文件
		f, err := os.Create("trace.out")
		if err != nil {
			fmt.Println("os.create err:", err.Error())
		}
		defer f.Close()
		//启动trace文件
		err = trace.Start(f)
		if err != nil {
			fmt.Println("trace.Start err:", err.Error())
		}
		//停止
		trace.Stop()

	*/

}

func ff() {
	fmt.Println("hello goroutine")
}
