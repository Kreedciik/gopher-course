package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Like interface {
	ToggleLike(string, string) error
}

type LikeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) *LikeRepository {
	return &LikeRepository{
		db,
	}
}

var (
	likesTable = "likes"
)

func (l *LikeRepository) ToggleLike(postId, userId string) error {
	var exists bool
	query := fmt.Sprintf(`SELECT EXISTS(
	SELECT 1 FROM %s WHERE post_id = $1 AND user_id = $2
	)`, likesTable)

	row := l.db.QueryRow(query, postId, userId)
	if err := row.Scan(&exists); err != nil {
		return err
	}

	if exists {
		deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE post_id = $1 AND user_id = $2`, likesTable)
		_, err := l.db.Exec(deleteQuery, postId, userId)
		if err != nil {
			return err
		}

		return nil
	}

	insertQuery := fmt.Sprintf(
		`INSERT INTO %s (id, post_id, user_id) VALUES ($1, $2, $3)`,
		likesTable)

	_, err := l.db.Exec(insertQuery, uuid.NewString(), postId, userId)
	return err
}
