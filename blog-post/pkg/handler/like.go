package handler

import (
	"blogpost/models"
	"blogpost/pkg/responses"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) likePost(ctx *gin.Context) {
	var likeDto models.LikeDTO
	value, _ := ctx.Get("userId")
	userId := value.(string)

	fmt.Println("ID: ", userId)
	if err := ctx.ShouldBindJSON(&likeDto); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Like.ReactToPost(likeDto.PostId, userId); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponse(ctx, "success")

}
