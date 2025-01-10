package repository

import (
	"blogpost/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Post interface {
	FindAllPosts(string) ([]models.Post, error)
	InsertPost(models.PostCreateDTO) error
	UpdatePost(models.PostUpdateDTO) error
	DeletePost(string) error
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db,
	}
}

var (
	postsTable = "posts"
)

func (p *PostRepository) FindAllPosts(authorId string) ([]models.Post, error) {

	var (
		posts []models.Post = []models.Post{}
		args  []interface{}
	)
	query := fmt.Sprintf(`
	SELECT p.id, p.title, p.description, u.id, u.name, u.email, 
	(SELECT COUNT(*) FROM likes WHERE likes.post_id = p.id) AS likes,
	(SELECT COUNT(*) FROM comments WHERE comments.post_id = p.id) AS comments,
	p.created_at
	FROM %s p
	INNER JOIN %s u ON p.user_id = u.id
	`, postsTable, usersTable)

	if authorId != "" {
		query = fmt.Sprintf("%s WHERE p.user_id = $1", query)
		args = append(args, authorId)
	}
	query = fmt.Sprintf("%s ORDER BY p.created_at DESC", query)

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Author.Id,
			&post.Author.Name,
			&post.Author.Email,
			&post.Likes,
			&post.Comments,
			&post.CreatedAt,
		)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
func (p *PostRepository) InsertPost(newPost models.PostCreateDTO) error {
	id := uuid.NewString()
	query := fmt.Sprintf(
		`INSERT INTO %s (id, title, description, user_id)
		VALUES ($1, $2, $3, $4)`, postsTable)
	_, err := p.db.Exec(
		query,
		id,
		newPost.Title,
		newPost.Description,
		newPost.UserId,
	)
	return err
}

func (p *PostRepository) UpdatePost(post models.PostUpdateDTO) error {
	query := fmt.Sprintf(
		`UPDATE %s SET title = $1, description = $2, updated_at = $3
		WHERE id = $4
		`, postsTable)
	_, err := p.db.Exec(
		query,
		post.Title,
		post.Description,
		time.Now(),
		post.Id,
	)
	return err
}

func (p *PostRepository) DeletePost(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, postsTable)
	_, err := p.db.Exec(query, id)
	return err
}
