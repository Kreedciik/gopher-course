package handler

import (
	"farmish/model"
	"farmish/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		New Farm
// @Description	Create a new farm
// @ID				farm-create
// @Tags			Farm
// @Accept			json
// @Produce		json
// @Param			input	body		model.CreateFarmDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/farm/create [post]
func (h *Handler) CreateFarm(ctx *gin.Context) {
	var newFarm model.CreateFarmDTO
	var reply string
	if err := ctx.ShouldBindJSON(&newFarm); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.farmClientRPC.Call("Handler.CreateFarm", newFarm, &reply); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, reply)
}

// @Summary		Update Farm
// @Description	Update a farm
// @ID				farm-update
// @Tags			Farm
// @Accept			json
// @Produce		json
// @Param			input	body		model.UpdateFarmDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/farm/update [put]
func (h *Handler) UpdateFarm(ctx *gin.Context) {
	var farm model.UpdateFarmDTO
	var reply string
	if err := ctx.ShouldBindJSON(&farm); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.farmClientRPC.Call("Handler.UpdateFarm", farm, &reply); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, reply)
}

// @Summary		Get farm
// @Description	Get farm by ID
// @ID				farm-get
// @Tags			Farm
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id"	"id"
// @Success		200		{object}	model.HTTPDataSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/farm/{id} [get]
func (h *Handler) GetFarm(ctx *gin.Context) {
	var reply model.Farm
	id, _ := ctx.Params.Get("id")
	err := h.farmClientRPC.Call("Handler.GetFarm", id, &reply)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponseWithData(ctx, gin.H{
		"data": reply,
	})
}

// @Summary		Delete Farm
// @Description	Delete a farm
// @ID				farm-delete
// @Tags			Farm
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id"	"id"	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/farm/{id} [delete]
func (h *Handler) DeleteFarm(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	var reply string
	if err := h.farmClientRPC.Call("Handler.DeleteFarm", id, &reply); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, reply)
}
