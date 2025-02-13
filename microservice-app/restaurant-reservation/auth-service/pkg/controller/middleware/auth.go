package middleware

import (
	helper "auth/pkg/helpers"
	"auth/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func AuthMiddleware(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)

	if header == "" {
		response.NewErrorResponse(ctx, http.StatusUnauthorized, "Authorization header is empty")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		response.NewErrorResponse(ctx, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	if len(headerParts[1]) == 0 {
		response.NewErrorResponse(ctx, http.StatusUnauthorized, "token is empty")
		return
	}

	claims, err := helper.ValidateToken(headerParts[1])
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set("userId", claims.Id)
	ctx.Next()
}
