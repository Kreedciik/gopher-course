package repository

import "database/sql"

type Repository struct {
	User    *UserRepository
	Post    *PostRepository
	Like    *LikeRepository
	Comment *CommentRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Post:    NewPostRepository(db),
		Like:    NewLikeRepository(db),
		Comment: NewCommentRepository(db),
	}
}
