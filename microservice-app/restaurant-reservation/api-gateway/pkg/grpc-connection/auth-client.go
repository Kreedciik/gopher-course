package grpcconnection

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthConnection(port string) (*grpc.ClientConn, error) {
	return grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
