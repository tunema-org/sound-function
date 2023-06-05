package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CategoryInput struct {
	Tag_ID int
}

// error kykny
func (h *handler) ListSamples(c *gin.Context) {
	var input CategoryInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
		return
	}

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
