package main

import (
	"net/http"

	"example.com/blog/controller"
	"example.com/blog/setup"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := setup.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//Go router declaration
	r := gin.Default()
	r.Use(setup.CorsMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	postController := controller.NewPostController(db)

	// Get all todos
	r.GET("/todos", postController.GetTodos)
	r.Run()
}
