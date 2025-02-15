package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	pb "reservation/grpc_gen/auth"
	"reservation/model"
	"reservation/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(ctx *gin.Context) {
	val, isExist := ctx.Get("userId")
	if !isExist {
		msg := "userId does not exists in the context"
		slog.Error(fmt.Sprintf("profile: %s", msg))
		response.NewErrorResponse(ctx, http.StatusInternalServerError, msg)
		return
	}
	userId := val.(string)
	user, err := h.authClient.GetProfile(ctx, &pb.GetProfileRequest{
		Id: userId,
	})
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponseWithData(ctx, model.User{
		Id:    user.GetId(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	})
}
