package main

import (
	"fmt"
	"time"

	a1 "exam-1/assignment1"
	a2 "exam-1/assignment2"
	a3 "exam-1/assignment3"
	a4 "exam-1/assignment4"
	a5 "exam-1/assignment5"
	a6 "exam-1/assignment6"
	a7 "exam-1/assignment7"
)

const (
	fileName  = "assignment7/students.json"
	fileName2 = "assignment7/students2.json"
)

func main() {
	a1.GetValueType("Hello")
	fmt.Println()

	fmt.Println(a2.GetNumberInWord(9999))
	fmt.Println()

	fmt.Println("Max value: ", a3.SearchMaxValue([]int{8, 11, 15, 99, 1, 25}))
	fmt.Println()

	fmt.Println("Missing value: ", a4.FindMissingValue([]int{1, 2, 3, 4, 6, 7, 8}))
	fmt.Println()

	a5.IncDecOperation([]string{"--X", "X++", "X++"})
	fmt.Println()

	ch := make(chan time.Time)
	go a6.TimeDistance(ch, time.Now().Add(2*time.Hour))
	ch <- time.Now()
	diff := <-ch
	fmt.Println(diff.UTC())
	fmt.Println()

	students, _ := a7.ReadStudentsFromFile(fileName)
	_ = a7.WriteStudentsToFile(students, fileName2)
	a7.PrintStudentDetails(students)
	topStudent, _ := a7.FindTopScoringStudent(students)
	fmt.Println("TOP: ", topStudent.Name)
	grouped, _ := a7.GroupStudentsByCategory(students, "CUJQWB")
	fmt.Println("The number of students: ", len(grouped["CUJQWB"]))
}
