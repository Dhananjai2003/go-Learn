package handler

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"url-shortner/db"
)

func Redirect(c *gin.Context) {

	code := c.Param("hash")

	var Url string

	err := db.DB.QueryRow("SELECT url FROM urls where hash = $1",code).Scan(&Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error":"We have trouble getting to our servers"})
		return
	}

	c.Redirect(http.StatusFound,Url)

}