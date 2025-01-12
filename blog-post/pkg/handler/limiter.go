package handler

import (
	"blogpost/pkg/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RateLimiter(ctx *gin.Context) {
	value, _ := ctx.Get("userId")
	userId := value.(string)

	if err := h.services.Limiter.CheckRequest(userId, ctx); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.Next()
}
