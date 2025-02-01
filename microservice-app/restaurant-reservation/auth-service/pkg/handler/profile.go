package handler

import (
	"auth/pkg/response"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(ctx *gin.Context) {
	val, isExist := ctx.Get("userId")
	if !isExist {
		msg := "userId does not exists in the context"
		slog.Error(fmt.Sprintf("profile: %s", msg))
		response.NewErrorResponse(ctx, http.StatusInternalServerError, msg)
		return
	}
	userId := val.(string)
	user, err := h.services.User.GetUserProfile(userId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponseWithData(ctx, user)
}
