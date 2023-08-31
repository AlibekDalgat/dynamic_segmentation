package handler

import (
	_ "github.com/AlibekDalgat/dynamic_segmentation/docs"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {

	h.DeleteExpirated()
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		segment := api.Group("segment")
		{
			segment.POST("/", h.createSegment)
			segment.POST("/:per", h.createSegmentWihtercent)
			segment.DELETE("/", h.deleteSegment)
		}

		user := api.Group("user")
		{
			user.PUT("/", h.updateSegsToUser)
			user.GET("/:user_id", h.getActiveSegments)
			user.POST("/", h.getReport)
		}

	}
	return router
}
