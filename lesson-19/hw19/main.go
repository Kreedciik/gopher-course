package main

import (
	"hw19/postgres"
	"hw19/repository"
	"hw19/utils"
)

func main() {
	pgDB, err := postgres.InitializePostgres()
	if err != nil {
		panic(err)
	}
	defer pgDB.Close()
	// studentRepo := repository.CreateStudentRepository(pgDB)
	courseRepo := repository.CreateCourseRepository(pgDB)
	// tutorRepo := repository.CreateTutorRepository(pgDB)
	// groupRepo := repository.CreateGroupRepository(pgDB)

	// newCourse := model.Course{
	// 	Id:            uuid.V4(),
	// 	Name:          "Introduction to Physics",
	// 	TutorID:       "307b308b-9e21-4f61-abd6-245b52615d89",
	// 	StudentNumber: 0,
	// 	StartedAt:     time.Now().Add(time.Duration(time.Now().Day()) + 2),
	// }
	// err = courseRepo.CreateCourse(newCourse)
	// utils.HandleError(err)
	// tutor := model.Tutor{Id: uuid.V4(), Name: "Tamara", LastName: "Soltesz", Email: "tamara@example.com"}
	// err = tutorRepo.CreateTutor(tutor)
	// utils.HandleError(err)
	// studentRepo.GetAllStudents()
	// courseRepo.GetCourses()
	// tutorRepo.GetAllTutors()
	// groupRepo.GetAllGroups()

	// oldestStudent, err := studentRepo.GetOldestStudent()
	// utils.HandleError(err)
	// fmt.Println(oldestStudent)
	// youngestStudent, err := studentRepo.GetYoungestStudent()
	// utils.HandleError(err)
	// fmt.Println(youngestStudent)

	// newGroup := model.StudentGroup{
	// 	Id:           uuid.V4(),
	// 	Name:         "420-18 Physics",
	// 	CourseId:     "b1b2e07a-5a5b-4ff3-843a-424b66f09287",
	// 	StudentCount: 0,
	// }
	// err = groupRepo.CreateGroup(newGroup)
	// utils.HandleError(err)

	// bestStudentByGroup, err := studentRepo.GetBestStudentByGroup("72b47cd5-1a74-4c2e-968d-2dee275725c9")
	// utils.HandleError(err)
	// fmt.Println(bestStudentByGroup)

	// bestStudentByCourse, err := studentRepo.GetBestStudentByCourse("9dd023e2-5a72-4570-91a2-e916d00a274a")
	// utils.HandleError(err)
	// fmt.Println(bestStudentByCourse)

	// err = courseRepo.EnrollToCourse("89e7a6b2-2f4c-4538-b5e9-5e9013456b8f", "9dd023e2-5a72-4570-91a2-e916d00a274a")
	// utils.HandleError(err)
	// err = courseRepo.EnrollToCourse("f7d6b953-7e58-4c8b-b7e3-cf9143469f6c", "9dd023e2-5a72-4570-91a2-e916d00a274a")
	// utils.HandleError(err)

	err = courseRepo.ExcludeFromCourse("f7d6b953-7e58-4c8b-b7e3-cf9143469f6c")
	utils.HandleError(err)
}
