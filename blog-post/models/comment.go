package models

import "time"

type CommentCreateDTO struct {
	PostId  string `json:"postId" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CommentUpdateDTO struct {
	Id      string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type Comment struct {
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	PostId    string    `json:"postId"`
	Author    User      `json:"commentAuthor"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentsResponse struct {
	TotalPages  float64   `json:"totalPages"`
	CurrentPage int       `json:"currentPage"`
	Comments    []Comment `json:"comments"`
}
