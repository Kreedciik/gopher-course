package handler

import (
	"farmish/model"
	"farmish/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		New animal
// @Description	Create a new animal
// @ID				animal-create
// @Tags			Animal
// @Accept			json
// @Produce		json
// @Param			input	body		model.CreateAnimalDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/animal/create [post]
func (h *Handler) CreateAnimal(ctx *gin.Context) {
	var newAnimal model.CreateAnimalDTO
	if err := ctx.ShouldBindJSON(&newAnimal); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.farmClientRPC.Call("Handler.CreateAnimal", newAnimal, ""); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, "animal created successfully")
}

// @Summary		Update animal
// @Description	Update animal
// @ID				animal-update
// @Tags			Animal
// @Accept			json
// @Produce		json
// @Param			input	body		model.UpdateAnimalDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/animal/update [put]
func (h *Handler) UpdateAnimal(ctx *gin.Context) {
	var animal model.UpdateAnimalDTO
	if err := ctx.ShouldBindJSON(&animal); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.farmClientRPC.Call("Handler.UpdateAnimal", animal, ""); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, "animal updated successfully")
}

// @Summary		Delete animal
// @Description	Delete animal
// @ID				animal-delete
// @Tags			Animal
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id"	"id"	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/animal/update [put]
func (h *Handler) DeleteAnimal(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if err := h.farmClientRPC.Call("Handler.DeleteAnimal", id, ""); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponse(ctx, "deleted successfully")
}

// @Summary		Get animal
// @Description	Get animal by ID
// @ID				animal-get
// @Tags			Animal
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id"	"id"
// @Success		200		{object}	model.HTTPDataSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/animal/{id} [get]
func (h *Handler) GetAnimalById(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	var animal model.Animal
	err := h.farmClientRPC.Call("Handler.GetAnimalById", id, &animal)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponseWithData(ctx, gin.H{
		"data": animal,
	})
}

// @Summary		Toggle animal state
// @Description	Change state of health
// @ID				animal-health
// @Tags			Animal
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id"	"id"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/animal/healthy [post]
func (h *Handler) ToggleHealthyAnimal(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	var reply string
	if err := h.farmClientRPC.Call("Handler.ToggleHealth", id, &reply); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(ctx, reply)
}

// @Summary		Toggle animal state
// @Description	Change state of hunger
// @ID				animal-hungry
// @Tags			Animal
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id"	"id"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/api/v1/animal/healthy [post]
func (h *Handler) ToggleHungryAnimal(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	var reply string
	if err := h.farmClientRPC.Call("Handler.ToggleHunger", id, &reply); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(ctx, reply)
}

func (h *Handler) ListenNotifications() {
	// subscriber := h.ps.Subscribe("notifications")

	// go func() {
	// 	for message := range subscriber {
	// 		// Websocket might be implemented
	// 		fmt.Println("Notification Received:", message)
	// 	}
	// }()
}
