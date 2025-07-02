package handler

import (
	"url-shortner/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Shortner(c *gin.Context) {

	var req struct {
		Url  string
	}

	if err := c.BindJSON(&req); err != nil || req.Url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Input"})
		return
	}
	
	code := "testingCode"

	_, err := db.DB.Exec("INSERT INTO urls (hash, url) values ($1, $2)",code,req.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"We are having issues with our internal servers"})
		return
	}	

	c.JSON(http.StatusOK,gin.H{"URL":"http:localhost:8080/"+code})

}