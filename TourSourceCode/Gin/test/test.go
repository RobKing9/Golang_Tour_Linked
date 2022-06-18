package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由
	r := gin.Default()

	//param和query参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(200, name+" is "+action)
	})
	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, fmt.Sprintf("%s is %s", name, role))

	})

	//表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		name := c.PostForm("username")
		psw := c.PostForm("password")
		//加密
		enpsw, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("bcrypt failed:", err.Error())
		}
		c.JSON(200, gin.H{
			"type": types,
			"name": name,
			"psw":  enpsw,
		})

	})

	//加载HTML模板
	r.LoadHTMLGlob("template/*")
	r.GET("/main", func(c *gin.Context) {
		c.HTML(200, "main.html", gin.H{"title": "测试"})
	})

	//重定向
	r.GET("/re", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/main")
	})

	//中间件
	r.GET("/middleware", MiddleWare(), func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("Req:", req)
		c.JSON(200, gin.H{"request": req})
	})

	//监听9999端口
	r.Run(":9999")
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("中间件开始执行...")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕！", status)
		t := time.Since(start)
		fmt.Println("Time:", t)
	}
}
