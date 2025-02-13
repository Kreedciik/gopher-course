package model

import "github.com/golang-jwt/jwt/v5"

type SignInDTO = CreateUserDTO

type Token struct {
	AccessToken string `json:"accessToken"`
}

type UserClaims struct {
	Id string `json:"userId"`
	jwt.RegisteredClaims
}
