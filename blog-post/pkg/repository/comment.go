package repository

import (
	"blogpost/models"
	"database/sql"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

type Comment interface {
	FindComments(page, size, postId string) (models.CommentsResponse, error)
	InsertComment(models.CommentCreateDTO, string) error
	UpdateComment(models.CommentUpdateDTO, string) error
	DeleteComment(string, string) error
}

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db,
	}
}

const commentsTable = "comments"

func (c *CommentRepository) FindComments(page, size int, postId string) (models.CommentsResponse, error) {
	var (
		comments []models.Comment = []models.Comment{}
		response models.CommentsResponse
	)

	query := fmt.Sprintf(
		`SELECT c.id, c.content, c.post_id, c.created_at,
		u.id, u.name, u.email
		FROM %s c
		INNER JOIN %s u ON c.user_id = u.id
		WHERE post_id = $1
		ORDER BY c.created_at
		`,
		commentsTable,
		usersTable,
	)

	countQuery := fmt.Sprintf(
		`SELECT COUNT(*) FROM %s WHERE post_id = $1`,
		commentsTable,
	)
	var totalCount int
	if err := c.db.QueryRow(countQuery, postId).Scan(&totalCount); err != nil {
		return response, err
	}

	rows, err := c.db.Query(query, postId)
	if err != nil {
		return response, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.Id,
			&comment.Content,
			&comment.PostId,
			&comment.CreatedAt,
			&comment.Author.Id,
			&comment.Author.Name,
			&comment.Author.Email,
		)

		if err != nil {
			return response, err
		}

		comments = append(comments, comment)
	}

	response.Comments = comments
	response.TotalPages = math.Ceil(float64(totalCount) / float64(size))
	response.CurrentPage = page

	return response, nil

}

func (c *CommentRepository) InsertComment(newComment models.CommentCreateDTO, userId string) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (id, content, post_id, user_id)
			VALUES ($1, $2, $3, $4)
		`,
		commentsTable,
	)

	_, err := c.db.Exec(query,
		uuid.NewString(),
		newComment.Content,
		newComment.PostId,
		userId,
	)
	return err
}

func (c *CommentRepository) UpdateComment(comment models.CommentUpdateDTO, userID string) error {
	query := fmt.Sprintf(
		`UPDATE %s SET content = $1, updated_at = $2 WHERE id = $3 AND user_id = $4`,
		commentsTable)
	_, err := c.db.Exec(query,
		comment.Content,
		time.Now(),
		comment.Id,
		userID,
	)
	return err
}

func (c *CommentRepository) DeleteComment(id string, userId string) error {
	query := fmt.Sprintf(
		`DELETE FROM %s WHERE id = $1 AND user_id = $2`,
		commentsTable,
	)
	_, err := c.db.Exec(query, id, userId)
	return err
}
