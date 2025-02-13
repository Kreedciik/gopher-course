package controller

import (
	_ "auth/docs"
	pb "auth/grpc_gen/auth"
	"auth/pkg/service"
	"google.golang.org/grpc"
)

type Controller struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Controller {
	return &Controller{
		services,
	}
}

func (h *Controller) InitServers(server *grpc.Server) {
	authServer := NewAuthServer(h.services.User)

	pb.RegisterAuthServiceServer(server, authServer)
}
