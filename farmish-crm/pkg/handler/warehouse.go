package handler

import (
	"farmish/model"
	"farmish/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		New stock
// @Description	Create a new stock
// @ID				stock-create
// @Tags			Warehouse
// @Accept			json
// @Produce		json
// @Param			input	body		model.CreateStockDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/warehouse/create/stock [post]
func (h *Handler) CreateStock(ctx *gin.Context) {
	var newStock model.CreateStockDTO
	if err := ctx.ShouldBindJSON(&newStock); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if newStock.Cost < 0 || newStock.Quantity < 0 {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "value should not be negative")
		return
	}

	if err := h.services.Warehouse.CreateStock(newStock); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, "stock created successfully")
}

// @Summary		Add feed
// @Description	Update a feed
// @ID				feed-update
// @Tags			Warehouse
// @Accept			json
// @Produce		json
// @Param			input	body		model.SupplyStockDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/warehouse/supply/feed [post]
func (h *Handler) SupplyFeedStock(ctx *gin.Context) {
	var stock model.SupplyStockDTO
	if err := ctx.ShouldBindJSON(&stock); err != nil {
		fmt.Println(err.Error())
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if stock.Quantity < 0 {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "value should not be negative")
		return
	}

	if err := h.services.Warehouse.SupplyFeedToWarehouse(stock); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, "stock supplied successfully")
}

// @Summary		Add medicine
// @Description	Update a medicine
// @ID				medicine-update
// @Tags			Warehouse
// @Accept			json
// @Produce		json
// @Param			input	body		model.SupplyStockDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/warehouse/supply/medicine [post]
func (h *Handler) SupplyMedicineStock(ctx *gin.Context) {
	var stock model.SupplyStockDTO
	if err := ctx.ShouldBindJSON(&stock); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if stock.Quantity < 0 {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "value should not be negative")
		return
	}

	if err := h.services.Warehouse.SupplyMedicineToWarehouse(stock); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, "stock supplied successfully")
}
