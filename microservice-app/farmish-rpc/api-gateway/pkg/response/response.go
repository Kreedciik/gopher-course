package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewErrorResponse(ctx *gin.Context, statusCode int, msg string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"error": msg,
		"code":  statusCode,
	})
}

func NewSuccessResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func NewSuccessResponseWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
