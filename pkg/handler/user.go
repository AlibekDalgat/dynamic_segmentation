package handler

import "github.com/gin-gonic/gin"

func (h *Handler) newSegsToUser(g *gin.Context) {

}
func (h *Handler) getActiveSegments(g *gin.Context) {

}

func (h *Handler) InitRoutes() *gin.Engine {
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
			user.PUT("/", h.newSegsToUser)
			user.GET("/:user_id", h.getActiveSegments)
		}

	}

	return router
}
