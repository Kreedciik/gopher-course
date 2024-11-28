package course

import (
	"fmt"
)

type CourseManager interface {
	AddCourse(Course) string
	RemoveCourse(int) string
	ListCourses() []string
}

type Course struct {
	CourseID   int
	CourseName string
}
type ProgrammingCourse struct {
	Category string
	Courses  []Course
}

type DesignCourse struct {
	Category string
	Courses  []Course
}

type LanguageCourse struct {
	Category string
	Courses  []Course
}

func removeCourse(slice []Course, courseID int) ([]Course, string) {
	for i, v := range slice {
		if v.CourseID == courseID {
			return append(slice[:i], slice[i+1:]...), v.CourseName
		}
	}
	return slice, ""
}

func getCoursesName(slice []Course) []string {
	courseNames := []string{}
	for _, v := range slice {
		courseNames = append(courseNames, v.CourseName)
	}
	return courseNames
}

func (p *ProgrammingCourse) AddCourse(course Course) string {
	fmt.Printf("Added course %s into programming courses \n", course.CourseName)
	p.Courses = append(p.Courses, course)
	return course.CourseName
}
func (p *ProgrammingCourse) RemoveCourse(courseID int) string {
	fmt.Printf("Removed course with ID %d from programming courses \n", courseID)
	currentCourses, removedCourse := removeCourse(p.Courses, courseID)
	p.Courses = currentCourses
	return removedCourse
}
func (p *ProgrammingCourse) ListCourses() []string {
	fmt.Printf("List of programming courses \n")
	return getCoursesName(p.Courses)
}

func (d *DesignCourse) AddCourse(course Course) string {
	fmt.Printf("Added course %s into design courses \n", course.CourseName)
	d.Courses = append(d.Courses, course)
	return course.CourseName
}
func (d *DesignCourse) RemoveCourse(courseID int) string {
	fmt.Printf("Removed course with ID %d from programming courses \n", courseID)
	currentCourses, removedCourse := removeCourse(d.Courses, courseID)
	d.Courses = currentCourses
	return removedCourse
}
func (d *DesignCourse) ListCourses() []string {
	fmt.Printf("List of design courses \n")
	return getCoursesName(d.Courses)
}

func (l *LanguageCourse) AddCourse(course Course) string {
	fmt.Printf("Added course %s into design courses \n", course.CourseName)
	l.Courses = append(l.Courses, course)
	return course.CourseName
}
func (l *LanguageCourse) RemoveCourse(courseID int) string {
	fmt.Printf("Removed course with ID %d from programming courses \n", courseID)
	currentCourses, removedCourse := removeCourse(l.Courses, courseID)
	l.Courses = currentCourses
	return removedCourse
}
func (l *LanguageCourse) ListCourses() []string {
	fmt.Printf("List of language courses \n")
	return getCoursesName(l.Courses)
}

func AddAnyCourse(c CourseManager, course Course) {
	c.AddCourse(course)
}
func ShowAnyCourses(c []CourseManager) {
	for _, v := range c {
		v.ListCourses()
	}
}
