package main

import (
	"farmish"
	"farmish/config"
	clientrpc "farmish/pkg/client-rpc"
	"farmish/pkg/handler"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := config.InitConfig(); err != nil {
		slog.Error(err.Error())
	}

	authClientRPC, err := clientrpc.NewAuthClientRPC(viper.GetString("rpc.auth"))
	if err != nil {
		slog.Error(fmt.Sprintf("auth-rpc-client: %s", err.Error()))
	}
	farmClientRPC, err := clientrpc.NewFarmClientRPC(viper.GetString("rpc.farm"))
	if err != nil {
		slog.Error(fmt.Sprintf("farm-rpc-client: %s", err.Error()))
	}

	h := handler.NewHandler(authClientRPC, farmClientRPC)

	routes := h.InitRoutes()

	server := new(farmish.Server)
	err = server.Run(viper.GetString("server.port"), routes)

	if err != nil {
		slog.Error(fmt.Sprintf("error when initialize server: %s", err.Error()))
	}
}
