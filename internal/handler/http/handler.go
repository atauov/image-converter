package http

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/atauov/image-converter/internal/service"
	"github.com/atauov/image-converter/internal/worker"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services *service.Service
	cfg      *config.HTTPServer
	asynqC   *worker.Client
}

func NewHandler(services *service.Service, cfg *config.HTTPServer, worker *worker.Client) *Handlers {
	return &Handlers{
		services: services,
		cfg:      cfg,
		asynqC:   worker,
	}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/:id", h.uploadImage)
		api.PUT("/", h.changeImage)
		api.GET("/", h.getAllImages)
		api.GET("/:id", h.getByKey)
		api.DELETE("/link/", h.deleteByURL)
		api.DELETE("/:id", h.deleteByKey)
	}

	return router
}
