package handler

import (
	"net/http"
	pb "reservation/grpc_gen/auth"
	"reservation/model"
	"reservation/pkg/response"

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
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	_, err := h.authClient.SignUp(ctx, &pb.SignUpRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	})

	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(ctx, "user registered successfully")
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
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	r, err := h.authClient.SignIn(ctx, &pb.SignInRequest{
		Email:    credentials.Email,
		Password: credentials.Password,
	})

	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponseWithData(ctx, model.Token{
		AccessToken: r.GetToken(),
	})

}
