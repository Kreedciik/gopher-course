package handler

import (
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
	dashboard, err := h.services.Dashboard.GetDashboardData()
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponseWithData(ctx, gin.H{
		"data": dashboard,
	})
}
