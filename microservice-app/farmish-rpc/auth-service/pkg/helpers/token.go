package helper

import (
	"farmish/model"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(user model.User, ttl time.Duration) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, model.UserClaims{
		user.Id,
		jwt.RegisteredClaims{
			Issuer:    "farmish",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	token, err := claims.SignedString([]byte(os.Getenv("SIGNING_KEY")))

	return token, err
}

func ValidateToken(accessToken string) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return model.UserClaims{}, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*model.UserClaims)

	if ok {
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			return nil, fmt.Errorf("token has expired")
		}

		if claims.Id == "" {
			return nil, fmt.Errorf("user ID is missing in token")
		}
	}
	return claims, nil
}
