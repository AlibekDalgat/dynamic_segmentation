package handler

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createSegment(c *gin.Context) {
	var input dynamic_segmentation.SegmentInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Segment.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) deleteSegment(c *gin.Context) {
	var input dynamic_segmentation.SegmentInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Segment.DeleteSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) createSegmentWihtercent(c *gin.Context) {
	percent, err := strconv.Atoi(c.Param("per"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неправильно введён процент")
		return
	}
	var input dynamic_segmentation.SegmentInfo
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateSegmentWihtPercent(percent, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
