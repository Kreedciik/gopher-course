package service

import (
	"blogpost/pkg/repository"
)

type Service struct {
	User
	Post
	Like
	Comment
}

func NewServices(repositories *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repositories.User),
		Post:    NewPostService(repositories.Post),
		Like:    NewLikeService(repositories.Like),
		Comment: NewCommentService(repositories.Comment),
	}
}
