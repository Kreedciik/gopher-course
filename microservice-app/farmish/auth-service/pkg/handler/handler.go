package handler

import (
	_ "farmish/docs"
	"farmish/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

// @title           Farmish API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @host      localhost:8080
// @BasePath  /api/v1
func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	return router
}
