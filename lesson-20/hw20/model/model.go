package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	StudentID string
	Name      string
	LastName  string
	Phone     string
	Age       int
	Grade     int
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Course struct {
	CourseId      string
	Name          string
	StudentNumber int
	TutorID       string
	StartedAt     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
type Tutor struct {
	Id        string
	Name      string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Group struct {
	Id           string
	Name         string
	CourseId     string
	StudentCount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Enroll struct {
	gorm.Model
	Id        string
	CourseID  string
	StudentID string
}
