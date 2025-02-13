package app

import (
	"auth/pkg/controller"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

type Server struct {
	RPC *controller.AuthController
}

func (s *Server) Run() error {

	log.Printf("server listening at %v", l.Addr())
	if err := server.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
