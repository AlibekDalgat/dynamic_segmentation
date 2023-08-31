package handler

import (
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Update user
// @Tags users
// @Description update inof about segments of user
// @ID update-segments-for-user
// @Accept  json
// @Produce  json
// @Param input body dynamic_segmentation.UserUpdatesInfo true "user update info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/ [put]
func (h *Handler) updateSegsToUser(c *gin.Context) {
	var input dynamic_segmentation.UserUpdatesInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
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

// @Summary Get active segments
// @Tags users
// @Description get active segments of user
// @ID get-active-segments
// @Produce  json
// @Success 200 {object} []dynamic_segmentation.SegmentInfo
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/:user_id [get]
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

// @Summary Get report
// @Tags users
// @Description get report
// @ID get-report
// @Accept  json
// @Produce  json
// @Success 200 {object} referenceFile
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/ [post]
func (h *Handler) getReport(c *gin.Context) {
	var input dynamic_segmentation.DateInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	file, err := h.services.GetReport(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, referenceFile{
		Reference: file.Name(),
	})
}
