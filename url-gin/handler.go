package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Url struct {
	Id int
	Name string
	Url string
}

var nextid = 1

var urlList []Url

func List(c *gin.Context){
	c.JSON(http.StatusOK,urlList)
}

func AddUrl(c *gin.Context){
	var currentUrl Url

	if err := c.BindJSON(&currentUrl); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Bad call"})
		return
	}

	currentUrl.Id = nextid
	nextid++

	urlList = append(urlList,currentUrl)
	c.JSON(http.StatusCreated,currentUrl)
}

func DeleteUrl(c *gin.Context){
	idStr := c.Param("id")
	id , err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid int"})
		return
	}

	for i, mess := range urlList {
		if mess.Id == id {
			urlList = append(urlList[:i],urlList[i+1:]...)
			c.JSON(http.StatusOK,gin.H{"Message":"Done"})
			return
		}
	}

	c.JSON(http.StatusBadRequest,gin.H{"Error:":"No message like that"})
}