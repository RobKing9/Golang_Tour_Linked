package main

import (
	"fmt"
	"github.com/kubastick/ginception"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//中间件
func MiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件MiddleWare1开始执行")
		c.Next()
		fmt.Println("中间件MiddleWare1执行完毕")
		t1 := time.Since(t)
		fmt.Printf("耗时%v", t1)

	}
}

func main() {
	//创建路由
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLFiles("./log.html")

	//注册路由

	//路由组blog
	blog := r.Group("/blog")
	{
		blog.GET("/article", MiddleWare1(), article)
		blog.GET("/image", image())
		//路由组嵌套
		//中间件hydra

		log := blog.Group("/log")
		{
			log.GET("/admin", ginception.Middleware(), func(c *gin.Context) {
				//c.JSON(http.StatusOK, gin.H{
				//	"status" : "ok",
				//})
				c.HTML(http.StatusOK, "log.html", gin.H{})
			})

			log.POST("/doadmin", doadmin())
		}
	}
	//路由组bili
	bili := r.Group("/bili")
	{
		bili.GET("/site", MiddleWare1(), func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "https://umei.cc/meinvtupian/xingganmeinv/")
		})
	}

	//启动路由
	r.Run(":9999")
}

func article(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "你访问的是文章页面",
	})
}

func image() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "你访问的是图片页面",
		})
	}
}

func doadmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//表单
		tel := c.PostForm("tel")
		psw := c.PostForm("psw")

		c.JSON(http.StatusOK, gin.H{
			"电话号码": tel,
			"密码":   psw,
		})

		fmt.Println()
		//文件
		form, _ := c.MultipartForm()
		//获取图片
		photo := form.File["files"]
		//遍历所有图片
		for _, file := range photo {
			//保存图片
			c.SaveUploadedFile(file, file.Filename)
			fmt.Println("图片保存成功！")
		}
		c.String(http.StatusOK, "登录成功！")

	}
}
