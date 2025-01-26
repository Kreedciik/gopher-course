package handler

import (
	"farmish/model"
	"farmish/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		New feeding
// @Description	Feeding of animal
// @ID				feeding-create
// @Tags			Animal
// @Accept			json
// @Produce		json
// @Param			input	body		model.CreateFeedingDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/animal/feed [post]
func (h *Handler) Feeding(ctx *gin.Context) {
	var feeding model.CreateFeedingDTO
	if err := ctx.ShouldBindJSON(&feeding); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.services.Feeding.FeedAnimal(feeding); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, "success")
}
