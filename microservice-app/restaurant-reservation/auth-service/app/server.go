package app

import (
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	Server   *grpc.Server
	Listener net.Listener
}

func NewGRPC(port string) (*GRPCServer, error) {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	return &GRPCServer{
		Server:   grpcServer,
		Listener: l,
	}, nil
}

func (s *GRPCServer) RunGRPC() error {
	return s.Server.Serve(s.Listener)
}
