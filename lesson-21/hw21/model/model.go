package model

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Filter struct {
	Pagination string `json:"pagination"`
	Size       string `json:"size"`
	Search     string `json:"search"`
	Users      []User `json:"users"`
}
