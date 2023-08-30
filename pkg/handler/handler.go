package handler

import (
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	h.DeleteExpirated()
	router := gin.New()
	api := router.Group("/api")
	{
		segment := api.Group("segment")
		{
			segment.POST("/", h.createSegment)
			segment.DELETE("/", h.deleteSegment)
		}

		user := api.Group("user")
		{
			user.PUT("/", h.updateSegsToUser)
			user.GET("/:user_id", h.getActiveSegments)
			user.GET("/", h.getReport)
		}

	}
	return router
}
