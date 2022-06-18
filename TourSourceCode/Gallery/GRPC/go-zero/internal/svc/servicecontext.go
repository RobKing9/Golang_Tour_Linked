package svc

import (
	"os_bufio_ioutil_file/ACode/Golang/GRPC/go-zero/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
