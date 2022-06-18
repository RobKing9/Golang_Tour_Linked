package router

import (
	"Experiment1/controler"
	"Experiment1/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouters(r *gin.Engine) *gin.Engine {
	r.GET("/", controler.Login)
	r.POST("/logsuccess", controler.LoginSuccess)

	r.GET("/register", controler.Register)
	r.POST("/regsuccess", controler.RegSuccess)

	r.GET("/index", middleware.AuthMiddleware(), controler.Index)
	return r
}
