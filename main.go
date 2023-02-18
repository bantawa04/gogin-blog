package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Go router declaration
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	r.Run()
}
