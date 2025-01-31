package main

import (
	"farmish"
	"farmish/config"
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
	h := handler.NewHandler()

	routes := h.InitRoutes()

	server := new(farmish.Server)
	err := server.Run(viper.GetString("server.port"), routes)

	if err != nil {
		slog.Error(fmt.Sprintf("error when initialize server: %s", err.Error()))
	}
}
