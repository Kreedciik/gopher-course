package repository

import (
	"database/sql"
	"hw19/model"
)

type groupRepository struct {
	db *sql.DB
}

func CreateGroupRepository(db *sql.DB) groupRepository {
	return groupRepository{db}
}

func (g *groupRepository) CreateGroup(group model.StudentGroup) error {
	_, err := g.db.Exec(`INSERT INTO student_group (group_id, name, course_id, 
		student_count) VALUES (
			$1, $2, $3, $4
		)`,
		group.Id, group.Name,
		group.CourseId, group.StudentCount,
	)
	return err
}

func (g *groupRepository) GetGroup(id string) (model.StudentGroup, error) {
	var group model.StudentGroup
	row := g.db.QueryRow(`SELECT group_id, name, course_id, 
		student_count
		FROM student_group
		WHERE group_id = $1
	`, id)

	err := row.Scan(
		&group.Id, &group.Name,
		&group.CourseId, &group.StudentCount,
	)

	return group, err
}

func (g *groupRepository) GetAllGroups() ([]model.StudentGroup, error) {
	var groups []model.StudentGroup

	rows, err := g.db.Query(
		`SELECT group_id, name, course_id, 
		student_count
		FROM student_group`)

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

func (g *groupRepository) UpdateGroup(group model.StudentGroup) error {
	_, err := g.db.Exec(`UPDATE student_group SET
	name = $1, course_id = $2, 
	student_count = $3, updated_at = $4, 
	WHERE group_id = $5
	`, group.Name, group.CourseId,
		group.StudentCount, group.UpdatedAt, group.Id,
	)

	return err
}

func (g *groupRepository) DeleteGroup(id string) error {
	_, err := g.db.Exec(`DELETE FROM student_group WHERE group_id = $1`, id)
	return err
}

func (g *groupRepository) GetBiggestGroup() (model.StudentGroup, error) {
	var group model.StudentGroup
	row := g.db.QueryRow(`SELECT group_id, name, course_id, 
		student_count
		FROM student_group
		ORDER BY student_count
		LIMIT 1
	`)

	err := row.Scan(
		&group.Id, &group.Name,
		&group.CourseId, &group.StudentCount,
	)

	return group, err
}
