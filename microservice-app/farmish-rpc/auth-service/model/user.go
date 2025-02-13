package model

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
