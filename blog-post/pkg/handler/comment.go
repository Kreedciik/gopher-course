package handler

import (
	"blogpost/models"
	"blogpost/pkg/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getComments(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	sizeStr := ctx.DefaultQuery("size", "10")
	postId := ctx.DefaultQuery("postId", "")

	if postId == "" {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, "postId should not be empty")
		return
	}

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, "could not parse page string to int")
		return
	}

	size, err := strconv.Atoi(sizeStr)

	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, "could not parse size string to int")
		return
	}

	commentResponse, err := h.services.Comment.GetComments(page, size, postId)

	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponseWithData(ctx, commentResponse)
}

func (h *Handler) createComment(ctx *gin.Context) {
	var newComment models.CommentCreateDTO
	value, _ := ctx.Get("userId")
	userId := value.(string)

	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Comment.CreateNewComment(newComment, userId); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponse(ctx, "comment successfully created")
}

func (h *Handler) updateComment(ctx *gin.Context) {
	var updatedComment models.CommentUpdateDTO
	if err := ctx.ShouldBindJSON(&updatedComment); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	value, _ := ctx.Get("userId")
	userId := value.(string)
	if err := h.services.Comment.EditComment(updatedComment, userId); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponse(ctx, "comment successfully updated")
}

func (h *Handler) deleteComment(ctx *gin.Context) {
	commentId, _ := ctx.Params.Get("id")
	value, _ := ctx.Get("userId")
	userId := value.(string)
	if err := h.services.Comment.DeleteComment(commentId, userId); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponse(ctx, "comment successfully deleted")
}
