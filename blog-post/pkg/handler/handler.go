package handler

import (
	"blogpost/pkg/handler/middleware"
	"blogpost/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	v1 := router.Group("/api/v1", middleware.AuthMiddleware, h.RateLimiter)
	{
		user := v1.Group("/user")
		{
			user.GET("/:id", h.getUser)
			user.PUT("/edit", h.updateUser)
			user.POST("/subscribe", h.followUser)
		}

		post := v1.Group("/post")
		{
			post.GET("", h.getAllPosts)
			post.POST("/create", h.createPost)
			post.PUT("/update", h.updatePost)
			post.DELETE("/delete/:id", h.deletePost)
		}

		like := v1.Group("/like")
		{
			like.POST("", h.likePost)
		}

		comment := v1.Group("/comment")
		{
			comment.GET("", h.getComments)
			comment.POST("/create", h.createComment)
			comment.PUT("/update", h.updateComment)
			comment.DELETE("delete/:id")
		}
	}

	return router
}
