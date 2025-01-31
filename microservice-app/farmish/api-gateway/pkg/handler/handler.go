package handler

import (
	_ "farmish/docs"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

// @title           Farmish API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @host      localhost:8080
// @BasePath  /api/v1
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Any("/*any", h.ProxyHandler)
	return router
}
