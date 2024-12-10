package assignment7

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
	Address Address  `json:"address"`
}

type Course struct {
	CourseID  string `json:"courseId"`
	Grade     string `json:"grade"`
	Professor string `json:"professor"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// Information about student
// ID: id
// Name: name
// Courses:

func printStudent(student Student) {
	borderWidth := 30
	fmt.Println(strings.Repeat("=", borderWidth))
	fmt.Printf("Student ID: %s\n", student.ID)
	fmt.Printf("Name      : %s\n", student.Name)
	fmt.Println("\nCourses:")
	fmt.Println(strings.Repeat("-", borderWidth))
	for _, course := range student.Courses {
		fmt.Printf("  Course ID : %s\n", course.CourseID)
		fmt.Printf("  Grade     : %s\n", course.Grade)
		fmt.Printf("  Professor : %s\n", course.Professor)
		fmt.Println(strings.Repeat("-", borderWidth))
	}
	fmt.Println("\nAddress:")
	fmt.Printf("  Street  : %s\n", student.Address.Street)
	fmt.Printf("  City    : %s\n", student.Address.City)
	fmt.Printf("  Country : %s\n", student.Address.Country)
	fmt.Println(strings.Repeat("=", borderWidth))
}

func WriteStudentsToFile(students []Student, filename string) error {
	jsonData, e := json.Marshal(students)
	if e != nil {
		return e
	}
	err := os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadStudentsFromFile(filename string) ([]Student, error) {
	bytes, err := os.ReadFile(filename)
	var students []Student
	if err != nil {
		return students, err
	}
	e := json.Unmarshal(bytes, &students)
	if e != nil {
		return students, e
	}
	return students, nil
}

func PrintStudentDetails(students []Student) {
	for _, student := range students {
		printStudent(student)
	}
}

func FindTopScoringStudent(students []Student) (*Student, error) {
	gradeScale := map[string]int{
		"A": 5,
		"B": 4,
		"C": 3,
		"D": 2,
		"E": 1,
		"F": 0,
	}
	topScored := struct {
		Max   int
		Index int
	}{Max: 0, Index: 0}
	for studentIndex, student := range students {
		maxScore := 0
		for _, course := range student.Courses {
			maxScore += gradeScale[course.Grade]
		}
		if topScored.Max < maxScore {
			topScored.Max = maxScore
			topScored.Index = studentIndex
		}
	}
	return &students[topScored.Index], nil
}

func GroupStudentsByCategory(students []Student, courseId string) (map[string][]Student, error) {
	grouped := map[string][]Student{courseId: []Student{}}

	for _, student := range students {
		for _, course := range student.Courses {
			if course.CourseID == courseId {
				grouped[courseId] = append(grouped[courseId], student)
			}
		}
	}
	return grouped, nil
}
