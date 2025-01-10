package handler

import (
	"blogpost/models"
	"blogpost/pkg/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUser(ctx *gin.Context) {
	userId, _ := ctx.Params.Get("id")
	user, err := h.services.User.GetUserById(userId)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponseWithData(ctx, gin.H{
		"data": user,
	})
}
func (h *Handler) updateUser(ctx *gin.Context) {
	var updatedUser models.UpdateUserDTO

	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.UpdateUser(updatedUser); err != nil {
		if err.Error() == "user is not found" {
			responses.NewErrorResponse(ctx, http.StatusNotFound, err.Error())
		} else {
			responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	responses.NewSuccessResponse(ctx, "successfully updated")

}

func (h *Handler) followUser(ctx *gin.Context) {
	value, _ := ctx.Get("userId")
	userId := value.(string)
	followerId := ctx.DefaultQuery("userId", "")

	if followerId == "" {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, "userId should not be empty")
		return
	}

	if followerId == userId {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, "user cannot subscribe himself")
		return
	}

	if err := h.services.User.FollowUser(userId, followerId); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	responses.NewSuccessResponse(ctx, "success")
}
