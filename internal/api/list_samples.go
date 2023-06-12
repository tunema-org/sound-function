package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) ListSamples(c *gin.Context) {
	samples, err := h.backend.ListSamples(c.Request.Context())
	if err != nil {
		log.Err(err).Msg("problem with creating sample")
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
