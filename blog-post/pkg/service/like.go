package service

import (
	"blogpost/pkg/repository"
)

type Like interface {
	ReactToPost(string, string) error
}

type LikeService struct {
	repository *repository.LikeRepository
}

func NewLikeService(repository *repository.LikeRepository) *LikeService {
	return &LikeService{
		repository: repository,
	}
}

func (l *LikeService) ReactToPost(postId, userId string) error {
	return l.repository.ToggleLike(postId, userId)
}
