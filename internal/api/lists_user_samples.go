package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListUserSamples(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid user id",
		})
		return
	}

	samples, err := h.backend.ListUserSamples(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, M{
		"message":     "sample retrieved",
		"items":       samples,
		"total_items": len(samples),
	})

}
