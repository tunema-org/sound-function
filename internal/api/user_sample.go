package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) UserSamples(c *gin.Context) {
	var input CategoryInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
		return
	}

	samples, err := h.backend.UserSamples(c.Request.Context())

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, M{
		"message":     "sample retrieved",
		"items":       samples,
		"total_items": len(samples),
	})

}
