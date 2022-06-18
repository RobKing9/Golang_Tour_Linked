package main

import (
	"context"
	"fmt"
	test "gin/proto"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
)

func CallRemote(c *gin.Context) {
	consulReg := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReg),
	)
	//初始化服务发现
	microClient := test.NewTestService("test", service.Client())
	//调用远程服务
	resp, err := microClient.Call(context.TODO(), &test.Request{
		Name:"zxw",
	})
	if err != nil {
		fmt.Println("call failed, err = ", err.Error())
		return
	}
	c.Writer.WriteString(resp.Msg)
	fmt.Println(resp, err)
}

func main () {

	route := gin.Default()
	route.GET("/", CallRemote)
	route.Run(":9999")
}
