package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) UserSamples(c *gin.Context) {
	samples, err := h.backend.UserSamples(c.Request.Context())

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
