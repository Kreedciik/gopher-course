package model

type User struct {
	Id       string `json:"id"`
	Username string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type CreateUserDTO struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
