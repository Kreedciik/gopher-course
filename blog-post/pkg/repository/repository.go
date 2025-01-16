package repository

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	User    *UserRepository
	Post    *PostRepository
	Like    *LikeRepository
	Comment *CommentRepository
}

func NewRepository(db *sql.DB, mongoDb *mongo.Database) *Repository {
	return &Repository{
		User:    NewUserRepository(db, mongoDb.Collection("users")),
		Post:    NewPostRepository(db),
		Like:    NewLikeRepository(db),
		Comment: NewCommentRepository(db),
	}
}
