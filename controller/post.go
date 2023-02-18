package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"example.com/blog/model"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	db *sql.DB
}

func NewPostController(db *sql.DB) *PostController {
	return &PostController{db}
}

func (pc *PostController) GetTodos(c *gin.Context) {
	var posts []model.Post
	rows, err := pc.db.Query(fmt.Sprintf("SELECT id, title FROM %s", model.Post{}.TableName()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.Title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, post)
	}
	c.JSON(http.StatusOK, posts)
}

// func (pc *PostController) GetTodo(c *gin.Context, id string) {
func (pc *PostController) GetTodo(c *gin.Context, id interface{}) {
	var post model.Post
	row := pc.db.QueryRow(fmt.Sprintf("SELECT id, title, content FROM %s WHERE id = %s", model.Post{}.TableName(), id))
	err := row.Scan(&post.ID, &post.Title, &post.Content)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
