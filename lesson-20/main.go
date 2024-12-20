package main

import (
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	connection := "host=localhost user=postgres password=root dbname=goonline port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&newUser)

	// newUser := User{Name: "Ilfat", Age: 26, Birthday: time.Now()}
	// result := db.Create(&newUser)
	// fmt.Println(result.RowsAffected)

	// Create many
	// users := []*User{
	// 	{Name: "Chris", Age: 20, Birthday: time.Now()},
	// 	{Name: "Agatha", Age: 30, Birthday: time.Now()},
	// 	{Name: "John", Age: 28, Birthday: time.Now()},
	// }
	// result := db.Create(users)
	// fmt.Println(result.RowsAffected)

	// Retrieve user
	// var users []User
	// db.Where("age IN ?", []int{20, 26}).Find(&users)
	// fmt.Println(users)

	//Update user
	// var user User

	// db.First(&user)
	// user.Name = "Updated name"
	// user.Age = 0
	// db.Save(&user)
	// db.Model(&User{}).Where("age IN ?", []int{20, 26}).Update("age", 10)

	// Delete user

	var user User
	db.Where("age = ?", 0).Delete(&user)
}
