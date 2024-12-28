package repository

import (
	"database/sql"
	"lesson22/model"

	"github.com/google/uuid"
)

type GroupRepository struct {
	db *sql.DB
}

func CreateGroupRepository(db *sql.DB) GroupRepository {
	return GroupRepository{db}
}

func (g *GroupRepository) CreateGroup(group model.StudentGroup) error {
	_, err := g.db.Exec(`INSERT INTO groups (group_id, name, course_id, 
		student_count) VALUES (
			$1, $2, $3, $4
		)`,
		uuid.New(), group.Name,
		group.CourseId, group.StudentCount,
	)
	return err
}

func (g *GroupRepository) GetGroup(id string) (model.StudentGroup, error) {
	var group model.StudentGroup
	row := g.db.QueryRow(`SELECT group_id, name, course_id, 
		groups
		FROM student_group
		WHERE group_id = $1
	`, id)

	err := row.Scan(
		&group.Id, &group.Name,
		&group.CourseId, &group.StudentCount,
	)

	return group, err
}

func (g *GroupRepository) GetAllGroups() ([]model.StudentGroup, error) {
	var groups []model.StudentGroup

	rows, err := g.db.Query(
		`SELECT group_id, name, course_id, 
		student_count
		FROM groups`)

	for rows.Next() {
		var group model.StudentGroup
		err = rows.Scan(
			&group.Id, &group.Name,
			&group.CourseId, &group.StudentCount,
		)
		groups = append(groups, group)
	}
	return groups, err
}

func (g *GroupRepository) UpdateGroup(group model.StudentGroup) error {
	_, err := g.db.Exec(`UPDATE groups SET
	name = $1, course_id = $2, 
	student_count = $3, updated_at = $4, 
	WHERE group_id = $5
	`, group.Name, group.CourseId,
		group.StudentCount, group.UpdatedAt, group.Id,
	)

	return err
}

func (g *GroupRepository) DeleteGroup(id string) error {
	_, err := g.db.Exec(`DELETE FROM groups WHERE group_id = $1`, id)
	return err
}

func (g *GroupRepository) GetBiggestGroup() (model.StudentGroup, error) {
	var group model.StudentGroup
	row := g.db.QueryRow(`SELECT group_id, name, course_id, 
		student_count
		FROM groups
		ORDER BY student_count
		LIMIT 1
	`)

	err := row.Scan(
		&group.Id, &group.Name,
		&group.CourseId, &group.StudentCount,
	)

	return group, err
}
