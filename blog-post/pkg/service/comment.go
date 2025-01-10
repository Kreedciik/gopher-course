package service

import (
	"blogpost/models"
	"blogpost/pkg/repository"
)

type Comment interface {
	GetComments(page, size int, postId string) (models.CommentsResponse, error)
	CreateNewComment(models.CommentCreateDTO, string) error
	EditComment(models.CommentUpdateDTO, string) error
	DeleteComment(string, string) error
}

type CommentService struct {
	repository *repository.CommentRepository
}

func NewCommentService(repository *repository.CommentRepository) *CommentService {
	return &CommentService{
		repository: repository,
	}
}

func (c *CommentService) GetComments(page, size int, postId string) (models.CommentsResponse, error) {
	return c.repository.FindComments(page, size, postId)
}

func (c *CommentService) CreateNewComment(newComment models.CommentCreateDTO, userId string) error {
	return c.repository.InsertComment(newComment, userId)
}

func (c *CommentService) EditComment(comment models.CommentUpdateDTO, userID string) error {
	return c.repository.UpdateComment(comment, userID)
}

func (c *CommentService) DeleteComment(id string, userId string) error {
	return c.repository.DeleteComment(id, userId)
}
