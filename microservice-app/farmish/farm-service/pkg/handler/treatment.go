package handler

import (
	"farmish/model"
	"farmish/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		New treatment
// @Description	Treatment of animal
// @ID				treatment-create
// @Tags			Animal
// @Accept			json
// @Produce		json
// @Param			input	body		model.CreateTreatmentDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/animal/treat [post]
func (h *Handler) Treatment(ctx *gin.Context) {
	var treatment model.CreateTreatmentDTO
	if err := ctx.ShouldBindJSON(&treatment); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.services.Treatment.TreatAnimal(treatment); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, "success")
}
