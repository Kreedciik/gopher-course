package middleware

import (
	"blogpost/pkg/helper"
	"blogpost/pkg/responses"
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
		responses.NewErrorResponse(ctx, http.StatusUnauthorized, "Athorization header is empty")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		responses.NewErrorResponse(ctx, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	if len(headerParts[1]) == 0 {
		responses.NewErrorResponse(ctx, http.StatusUnauthorized, "token is empty")
		return
	}

	claims, err := helper.ValidateToken(headerParts[1])
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set("userId", claims.Id)
	ctx.Next()
}
