package model

import "github.com/golang-jwt/jwt/v5"

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken string `json:"accessToken"`
}

type UserClaims struct {
	Id string `json:"userId"`
	jwt.RegisteredClaims
}
