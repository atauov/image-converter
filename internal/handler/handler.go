package handler

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/atauov/image-converter/internal/service"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services *service.Service
	cfg      *config.Config
}

func NewHandler(services *service.Service, cfg *config.Config) *Handlers {
	return &Handlers{
		services: services,
		cfg:      cfg,
	}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, "HELLO")
	})

	api := router.Group("/api")
	{
		api.POST("/:id", h.uploadImage)
		api.PUT("/:id", h.changeImage)
		api.GET("/", h.getAllImages)
		api.GET("/:id", h.getByKey)
		api.DELETE("/link/:id", h.deleteByURL)
		api.DELETE("/:id", h.deleteByKey)
	}

	return router
}
