package handler

import (
	"farmish/model"
	"farmish/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		User sign-up
// @Description	Create user
// @ID				sign-up
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			input	body		model.CreateUserDTO	true	"Body"
// @Success		200		{object}	model.HTTPSuccess
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/auth/sign-up [post]
func (h *Handler) SignUp(ctx *gin.Context) {
	var newUser model.CreateUserDTO
	var reply *string
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.authClientRPC.Call("Handler.SignUp", newUser, reply); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(ctx, *reply)
}

// @Summary		User sign-in
// @Description	Get access token
// @ID				sign-in
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			input	body		model.SignInDTO	true	"Credentials"
// @Success		200		{object}	model.Token
//
// @Failure		400		{object}	model.HTTPError
// @Failure		401		{object}	model.HTTPError
//
// @Router			/auth/sign-in [post]
func (h *Handler) SignIn(ctx *gin.Context) {
	var credentials model.SignInDTO
	var reply model.Token

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.authClientRPC.Call("Handler.SignIn", credentials, &reply); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.NewSuccessResponseWithData(ctx, reply)
}
