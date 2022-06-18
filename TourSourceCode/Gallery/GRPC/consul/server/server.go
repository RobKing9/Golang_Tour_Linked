package main

import (
	"consul/pb/person"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Children struct {
	person.UnimplementedSayHelloServer
}

// Greet 绑定类方法 实现接口
func (c *Children) Greet(ctx context.Context, p *person.Person) (*person.Person, error) {
	p.Name = fmt.Sprintf("hello %s", p.Name)
	return p, nil
}

func main() {
	//grpc注册到consul
	//1、初始化consul配置
	consulConfig := api.DefaultConfig()
	//2、创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	//3、告诉consul 即将注册的服务配置信息
	reg := api.AgentServiceRegistration{
		ID:      "1",
		Tags:    []string{"greet", "grpctest"},
		Name:    "HelloService",
		Address: "localhost",
		Port:    9999,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "localhost:9999",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	//4、注册服务到consul上
	consulClient.Agent().ServiceRegister(&reg)
	//1、初始化grpc对象
	grpcserver := grpc.NewServer()
	//2、注册服务
	person.RegisterSayHelloServer(grpcserver, &Children{})
	//3、设置监听
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("net.listen failed:", err.Error())
		return
	}
	defer listener.Close()
	grpcserver.Serve(listener)
	fmt.Println("服务启动中......")
	time.Sleep(time.Second)
	//4、启动服务
	fmt.Println("服务启动成功！！！\n正在监听9999端口！！！")
}
