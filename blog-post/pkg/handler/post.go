package handler

import (
	"blogpost/models"
	"blogpost/pkg/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllPosts(ctx *gin.Context) {
	authorId := ctx.DefaultQuery("authorId", "")
	posts, err := h.services.Post.GetAllPosts(authorId)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponseWithData(ctx, gin.H{
		"data": posts,
	})
}

func (h *Handler) createPost(ctx *gin.Context) {
	var newPost models.PostCreateDTO
	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Post.CreatePost(newPost); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.NewSuccessResponse(ctx, "post successfully created")
}
func (h *Handler) updatePost(ctx *gin.Context) {
	var updatedPost models.PostUpdateDTO
	if err := ctx.ShouldBindJSON(&updatedPost); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Post.UpdatePost(updatedPost); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	responses.NewSuccessResponse(ctx, "post successfully updated")
}
func (h *Handler) deletePost(ctx *gin.Context) {
	postId := ctx.Param("id")
	if postId == "" {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, "id is not provided")
		return
	}

	if err := h.services.Post.DeletePost(postId); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	responses.NewSuccessResponse(ctx, "post successfully deleted")
}
