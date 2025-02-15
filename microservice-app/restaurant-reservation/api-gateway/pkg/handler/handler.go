package handler

import (
	_ "reservation/docs"
	pb "reservation/grpc_gen/auth"
	"reservation/pkg/handler/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	authClient pb.AuthServiceClient
}

func NewHandler(authClient pb.AuthServiceClient) *Handler {
	return &Handler{
		authClient,
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
