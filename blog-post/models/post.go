package models

import "time"

type PostAuthor struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Post struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Author      PostAuthor `json:"author"`
	Likes       int        `json:"likes"`
	Comments    int        `json:"comments"`
	CreatedAt   time.Time  `json:"postedAt"`
}
type PostCreateDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserId      string `json:"userId" binding:"required"`
}

type PostUpdateDTO struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
