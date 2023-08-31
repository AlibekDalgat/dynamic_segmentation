package handler

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create segment
// @Tags segments
// @Description create segment
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body dynamic_segmentation.SegmentInfo true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/segment [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input dynamic_segmentation.SegmentInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
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

// @Summary Delete segment
// @Tags segments
// @Description delete segment
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param input body dynamic_segmentation.SegmentInfo true "segment info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/segment [delete]
func (h *Handler) deleteSegment(c *gin.Context) {
	var input dynamic_segmentation.SegmentInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
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

// @Summary Create segment
// @Tags segments
// @Description create segment
// @ID create-segment-with-percent
// @Accept  json
// @Produce  json
// @Param input body dynamic_segmentation.SegmentInfo true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/segment/:per [post]
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
