package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"
type User struct{
	Name string `form:"name" binding:"required,len=6"`
	Age int `form:"age" binding:"numeric,min=18,max=100"`
}

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":"admin",
			"age":21,
		})
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"name":"admin",
			"age":21,
		})
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static","./static")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
					"name":"admin",
					"age":21,
					"users": []string{"陈亚萌","张三","喜羊羊"},
		})
	})
	r.GET("/get", func(c *gin.Context) {
		fmt.Println(c.Query("name"))
		fmt.Println(c.Query("age"))
		c.String(http.StatusOK, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		fmt.Println(c.PostForm("name"))
		fmt.Println(c.PostForm("age"))
		c.String(http.StatusOK, "post")
	})
	r.GET("/get_user/:id", func(c *gin.Context) {
		fmt.Println(c.Param("id"))

		c.String(http.StatusOK,"url param")
	})
	r.GET("/getUser", func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)
		if err!=nil{
			c.String(http.StatusOK, err.Error())
		}else{
			c.String(http.StatusOK, "name->%s age -> %d", user.Name, user.Age)
		}
	})
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK,"upload.html",nil)
	})
	//上传单个文件
	r.POST("/upload", func(c *gin.Context) {
		//f, err := c.FormFile("file")
		//if err!=nil{
		//	log.Println(err)
		//}
		//err = c.SaveUploadedFile(f, fmt.Sprintf("upload/%s", f.Filename))
		//if err != nil{
		//	c.String(http.StatusOK,"上传文件失败%v", err.Error())
		//
		//}else{
		//	c.String(http.StatusOK, "上传成功！")
		//}
		form,_:=c.MultipartForm()
		files := form.File["file"]
		for _, file := range files{
			c.SaveUploadedFile(file, fmt.Sprintf("upload/%s", file.Filename))
		}
		c.String(http.StatusOK, "ok")
	})
	//上传多个文件


	//路由组
	api := r.Group("/api")
	{
		api.GET("/get_user", func(context *gin.Context) {
			context.String(http.StatusOK, "api get user")
		})
		api.GET("/get_info", func(context *gin.Context) {
			context.String(http.StatusOK, "api get info")
		})

	}

	admin := r.Group("/admin")
	{
		admin.GET("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "admin login")
		})
		admin.GET("/logout", func(context *gin.Context) {
			context.String(http.StatusOK, "admin logout")
		})
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
