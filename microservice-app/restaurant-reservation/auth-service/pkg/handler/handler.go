package handler

import (
	_ "auth/docs"
	"auth/pkg/handler/middleware"
	"auth/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

// @title           Auth service API
// @version         1.0
// @description     This is a sample server caller server.
// @name						Authorization
// @host      localhost:8081
// @BasePath  /api
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
		auth.GET("/profile", middleware.AuthMiddleware, h.GetProfile)
	}

	return router
}
