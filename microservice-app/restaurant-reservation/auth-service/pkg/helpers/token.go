package helper

import (
	"auth/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(user model.User, ttl time.Duration) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, model.UserClaims{
		user.Id,
		jwt.RegisteredClaims{
			Issuer:    "auth",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	token, err := claims.SignedString([]byte(os.Getenv("SIGNING_KEY")))

	return token, err
}
