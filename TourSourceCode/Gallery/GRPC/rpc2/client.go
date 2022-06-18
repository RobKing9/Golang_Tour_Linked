package main

import (
	"fmt"
	"net/rpc"
)

func main () {
	//连接服务器
	conn, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("dial failed, err = ", err.Error())
		return
	}
	defer conn.Close()
	//远程调用函数
	var resp string
	err = conn.Call("hello.HelloWorld", "zxw", &resp)
	if err != nil {
		fmt.Println("call failed, err = ", err.Error())
		return
	}
	fmt.Println(resp)
}
