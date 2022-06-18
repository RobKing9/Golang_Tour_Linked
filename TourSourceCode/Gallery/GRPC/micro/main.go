package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"renthome/service/Captcha/handler"
	proto "renthome/service/Captcha/proto"
)

func main() {
	//Consul Registry
	consulRegistry := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Registry(consulRegistry),
		micro.Address("127.0.0.1:10001"),
		micro.Name("captcha"),
		micro.Version("latest"),
	)

	// Register handler
	proto.RegisterCaptchaHandler(srv.Server(), new(handler.C))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
