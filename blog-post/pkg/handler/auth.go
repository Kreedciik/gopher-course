package handler

import (
	"blogpost/models"
	"blogpost/pkg/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var userCreateDTO models.UserCreateDTO
	if err := ctx.ShouldBindJSON(&userCreateDTO); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.CreateUser(userCreateDTO); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponse(ctx, "user successfully created")
}
func (h *Handler) signIn(ctx *gin.Context) {
	var credentials models.UserSignInDTO

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.User.LoginUser(credentials)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponseWithData(ctx, gin.H{
		"accessToken": token,
	})
}
