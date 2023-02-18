package models

type Post struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
