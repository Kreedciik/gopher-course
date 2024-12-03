package main

import (
	"encoding/json"
	"fmt"
)

var userJSON = `
{
"id": 1,
"firstName": "Emily",
"lastName": "Johnson",
"maidenName": "Smith",
"age": 28,
"gender": "female",
"email": "emily.johnson@x.dummyjson.com",
"phone": "+81 965-431-3024",
"hair": {
"color": "Brown",
"type": "Curly"
}
}`

type Hair struct {
	Color string
	Type  string
}
type User struct {
	Id         int
	FirstName  string
	LastName   string
	MaidenName string
	Age        int
	Gender     string
	Email      string
	Phone      string
	Hair       Hair
}

func main() {
	var user = User{}

	err := json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}
