package handler

import (
	"farmish/model"
	"farmish/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Dashboard
// @Description	Dashboard info
// @ID				dashboard-get
// @Tags			Dashboard
// @Accept			json
// @Produce		json
// @Success		200		{object}	model.HTTPDataSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/dashboard [get]
func (h *Handler) Dashboard(ctx *gin.Context) {
	var reply model.Dashboard
	if err := h.farmClientRPC.Call("Handler.Dashboard", "", &reply); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponseWithData(ctx, gin.H{
		"data": reply,
	})
}
