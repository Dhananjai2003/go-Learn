package main

import(
	"url-shortner/db"
	"url-shortner/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect()

	r := gin.Default()
	r.POST("/short",handler.Shortner)
	r.GET(":hash",handler.Redirect)

	r.Run(":8080")

}