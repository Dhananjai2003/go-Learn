package main

import (
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()

	r.GET("/list",List)

	r.POST("/add",AddUrl)

	r.DELETE("/delete/:id",DeleteUrl)

	r.Run(":8080")
}