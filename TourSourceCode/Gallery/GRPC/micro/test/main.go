package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"test/handler"
	pb"test/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	//初始化服务发现
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string {
			"localhost:9999",
		}
	})
	//添加服务
	srv := micro.NewService(
		micro.Name("test"),
		micro.Registry(consulReg),
		micro.Version("latest"),
		)
	// Create service
	//srv := service.New(
	//	service.Name("test"),
	//	service.Version("latest"),
	//)

	// Register handler
	pb.RegisterTestHandler(srv.Server(), new(handler.Test))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

//go get -u github.com/golang/protobuf/proto
//go get -u github.com/golang/protobuf/protoc-gen-go
//go get github.com/micro/micro/v3/cmd/protoc-gen-micro
