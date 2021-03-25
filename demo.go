package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func Test() gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL)
	}
}
func s() {
	r := gin.New()
	//使用中间件
	r.Use(Test())
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get method")
	})
	r.Run()
}
