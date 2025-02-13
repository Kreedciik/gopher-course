package handler

import (
	"farmish/pkg/pubsub"
	"farmish/pkg/service"
)

type Handler struct {
	services *service.Service
	ps       *pubsub.PubSub
}

func NewHandler(services *service.Service, ps *pubsub.PubSub) *Handler {
	return &Handler{
		services,
		ps,
	}
}
