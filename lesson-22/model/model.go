package model

import "time"

type Student struct {
	StudentID string    `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Age       int       `json:"age"`
	Grade     int       `json:"grade"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type Course struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	StudentNumber int       `json:"studentNumber"`
	TutorID       string    `json:"tutorId"`
	StartedAt     time.Time `json:"startedAt"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
type Tutor struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type StudentGroup struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	CourseId     string    `json:"courseId"`
	StudentCount int       `json:"studentCount"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
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

type StudentResponse struct {
	Data []Student `json:"data"`
}

type CourseResponse struct {
	Data []Course `json:"data"`
}

type TutorResponse struct {
	Data []Tutor `json:"data"`
}

type GroupResponse struct {
	Data []StudentGroup `json:"data"`
}
