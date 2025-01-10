package models

import "github.com/golang-jwt/jwt/v5"

type Follower struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type User struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Followers []Follower `json:"followers"`
}
type UserCreateDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdateUserDTO struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type UserSignInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserClaims struct {
	Id string `json:"userId"`
	jwt.RegisteredClaims
}
