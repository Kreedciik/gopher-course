package controller

import (
	pb "reservation/grpc_gen/reservation"
	"reservation/pkg/service"

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
	reservationServer := NewReservationServer(h.services)
	pb.RegisterReservationServiceServer(server, reservationServer)
}
