package model

import "time"

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
	Id            string
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
type StudentGroup struct {
	Id           string
	Name         string
	CourseId     string
	StudentCount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Enroll struct {
	CourseID  string
	StudentID string
}

type CreateStudentRequest struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Phone    string `json:"phone"`
	Age      int    `json:"age"`
	Grade    int    `json:"grade"`
	Gender   string `json:"gender"`
}
