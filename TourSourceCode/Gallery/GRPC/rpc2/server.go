package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct {

}

func(h *Hello) HelloWorld(name string, resp *string) error {
	*resp = name + "hello"
	return nil
}

func main () {
	//注册rpc服务
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println("register failed, err = ", err.Error())
		return
	}
	//监听端口
	listener, err := net.Listen("tcp", ":9999")
	fmt.Println("正在监听9999端口 ...")
	if err != nil {
		fmt.Println("listen failed, err = ", err.Error())
		return
	}
	defer listener.Close()
	//建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err = ", err.Error())
			return
		}
		defer conn.Close()

		fmt.Println("连接成功！！！")
		//绑定服务
		rpc.ServeConn(conn)
	}
}
