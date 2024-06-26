package handler

import (
	"github.com/atauov/image-converter/internal/service"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handlers {
	return &Handlers{
		services: services,
	}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/", h.uploadImage)
		api.PUT("/", h.changeImage)
		api.GET("/", h.getAllImages)
		api.GET("/:key", h.getByKey)
		api.DELETE("/:url", h.deleteByURL)
		api.DELETE("/:key", h.deleteByKey)
	}

	return router
}
