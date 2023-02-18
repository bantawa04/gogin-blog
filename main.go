package main

import (
	"net/http"

	"example.com/blog/models"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := models.ConnectToDB(); err != nil {
		panic(err)
	}
	//Go router declaration
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	r.Run()
}
