package main

import (
	"consul/pb/person"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	//consul
	//1、初始化consul配置
	consulConfig := api.DefaultConfig()
	//2、创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient failed:", err.Error())
	}
	//3、服务发现
	service, _, err := consulClient.Health().Service("HelloService", "grpctest", true, nil)
	//ip+port
	addr := service[0].Service.Address + ":" + strconv.Itoa(service[0].Service.Port)
	//1、连接服务
	//conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	//consul 方式连接服务
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial failed:", err.Error())
		return
	}
	//2、初始化客户端
	grpcClient := person.NewSayHelloClient(conn)
	//3、调用远程函数
	var person person.Person
	person.Name = "zxw"
	p, err := grpcClient.Greet(context.TODO(), &person)
	if err != nil {
		fmt.Println("greet failed:", err.Error())
		return
	}
	fmt.Println(p)
}
