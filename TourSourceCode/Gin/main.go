package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Run(":8000")

	/*
		上传文件
	*/
	/*
		r.LoadHTMLFiles("./template/file.html")
		r.GET("/file", func(c *gin.Context) {
			c.HTML(http.StatusOK, "file.html", nil)
		})
		r.POST("/upload", func(c *gin.Context) {
			//c.String(http.StatusOK, "上传成功！")
			//从请求中读取文件
			//单文件上传
			//file, err := c.FormFile("f")
			//if err != nil {
			//	c.JSON(http.StatusBadRequest, gin.H{
			//		"error" : error.Error(err),
			//	})
			//} else {//将文件保存到本地
			//	//in := fmt.Sprintf("/%s", file.Filename)
			//	c.SaveUploadedFile(file, file.Filename)
			//	c.JSON(http.StatusOK, gin.H{
			//		"msg":"上传成功",
			//	})
			//}
			//多文件上传
			form, err := c.MultipartForm()
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("err %s"), err.Error())
			}
			//获取所有文件
			files := form.File["files"]
			//遍历
			for _, file := range files {
				c.SaveUploadedFile(file, file.Filename)
				c.JSON(http.StatusOK, gin.H{
					"msg": "上传成功",
				})

			}
		})

	*/
}

/*
	获取请求的path URL参数
*/
//type User struct {
//	Username string `form:"username"`
//	Password string `form:"password"`
//}
//加载HTML模板
//r.LoadHTMLFiles("./template/Admin/login.html")
//r.GET("/:name/:age", func (c *gin.Context){
//	//获取路径参数
//	name := c.Param("name")
//	age := c.Param("age")
//	c.String(http.StatusOK, "姓名:%v\n年龄:%v \n", name, age)
//	c.JSON(http.StatusOK, gin.H{
//		"姓名":name,
//		"年龄":age,
//	})
//	c.HTML(http.StatusOK, "login.html", nil)
//
//})

//r.GET("/user", func(c *gin.Context) {
//	var u User
//	err := c.ShouldBind(&u)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error":err.Error(),
//		})
//	} else {
//		fmt.Printf("%#v\n", u)
//		c.JSON(http.StatusOK, gin.H{
//			"msg" : "okok",
//		})
//	}
//	//name := c.DefaultQuery("name", "zxw")
//	//c.JSON(http.StatusOK, gin.H{
//	//	"name":name,
//	//})
//})

///*
// 获取form表单数据
// */
////配置模板文件
////r.LoadHTMLGlob("template/**/*")
//r.LoadHTMLFiles("./template/Admin/login.html")
//r.GET("/admin", func(c *gin.Context) {
//	c.HTML(http.StatusOK, "login.html", gin.H{})
//})
//
//r.POST("/login", func(c *gin.Context) {
//
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	age := c.DefaultPostForm("age", "20")
//	c.JSON(http.StatusOK, gin.H{
//		"username" : username,
//		"password" : password,
//		"age" : age,
//	})
//})

////配置路由
//r.GET("/hello", func(c *gin.Context) {
//	c.String(http.StatusOK, "hello world hello gin")
//})
//r.GET("/json", func(c *gin.Context) {
//	c.JSONP(http.StatusOK, gin.H{
//		"success":true,
//		"title":"hello",
//	})
//})
//r.GET("/xml", func(c *gin.Context) {
//	c.XML(http.StatusOK, map[string]interface{}{
//		"success":true,
//		"title":"hello",
//	})
//})
//
//r.PUT("/edit", func(c *gin.Context) {
//	c.String(http.StatusOK, "put请求主要用于编辑数据")
//})
//r.DELETE("/delete", func(c *gin.Context) {
//	c.String(http.StatusOK, "Delete请求主要用于删除数据")
//})
