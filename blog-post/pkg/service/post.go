package service

import (
	"blogpost/models"
	"blogpost/pkg/repository"
)

type Post interface {
	GetAllPosts(string) ([]models.Post, error)
	CreatePost(models.PostCreateDTO) error
	UpdatePost(models.PostUpdateDTO) error
	DeletePost(string) error
}

type PostService struct {
	repository *repository.PostRepository
}

func NewPostService(postRepository *repository.PostRepository) *PostService {
	return &PostService{
		repository: postRepository,
	}
}

func (p *PostService) GetAllPosts(authorId string) ([]models.Post, error) {
	return p.repository.FindAllPosts(authorId)
}
func (p *PostService) CreatePost(newPost models.PostCreateDTO) error {
	return p.repository.InsertPost(newPost)
}

func (p *PostService) UpdatePost(post models.PostUpdateDTO) error {
	return p.repository.UpdatePost(post)
}

func (p *PostService) DeletePost(id string) error {
	return p.repository.DeletePost(id)
}
