package handler

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) updateSegsToUser(c *gin.Context) {
	var input dynamic_segmentation.UserUpdatesInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	errAdd, errDel := h.services.User.AddToSegments(input), h.services.User.DeleteFromSegments(input)
	if len(errAdd) != 0 || len(errDel) != 0 {
		for _, err := range errAdd {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		for _, err := range errDel {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getActiveSegments struct {
	Segments []dynamic_segmentation.SegmentInfo `json:"segments"`
}

func (h *Handler) getActiveSegments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неправильный id-параметр")
		return
	}
	segments, err := h.services.User.GetActiveSegments(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getActiveSegments{
		Segments: segments,
	})

}
