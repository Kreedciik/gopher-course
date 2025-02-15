package main

import (
	"fmt"
	"log/slog"
	"os"
	"reservation/app"
	"reservation/config"
	pb "reservation/grpc_gen/auth"
	grpcconnection "reservation/pkg/grpc-connection"
	"reservation/pkg/handler"

	"github.com/spf13/viper"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := config.InitConfig(); err != nil {
		slog.Error(err.Error())
	}

	authConn, err := grpcconnection.NewAuthConnection(viper.GetString("services.auth"))
	if err != nil {
		slog.Error(fmt.Sprintf("error when initialize auth connection: %s", err.Error()))
	}
	defer authConn.Close()
	authClient := pb.NewAuthServiceClient(authConn)

	h := handler.NewHandler(authClient)

	routes := h.InitRoutes()

	server := new(app.Server)
	err = server.Run(viper.GetString("server.port"), routes)

	if err != nil {
		slog.Error(fmt.Sprintf("error when initialize server: %s", err.Error()))
	}
}
