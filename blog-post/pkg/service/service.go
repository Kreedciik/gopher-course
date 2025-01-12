package service

import (
	"blogpost/pkg/repository"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	User
	Post
	Like
	Comment
	Limiter
}

func NewServices(repositories *repository.Repository, rdb *redis.Client) *Service {
	return &Service{
		User:    NewUserService(repositories.User),
		Post:    NewPostService(repositories.Post),
		Like:    NewLikeService(repositories.Like),
		Comment: NewCommentService(repositories.Comment),
		Limiter: NewLimiterService(rdb),
	}
}
