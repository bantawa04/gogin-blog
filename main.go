package main

import (
	"net/http"

	"example.com/blog/setup"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	if err := setup.ConnectToDB(); err != nil {
		panic(err)
	}
	//Go router declaration
	r := gin.Default()
	r.Use(setup.CorsMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	r.Run()
}
