package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) DeleteSample(c *gin.Context) {

	sampleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, M{
			"message": "invalid sample id",
		})
		return
	}

	err = h.backend.DeleteSample(c.Request.Context(), sampleID)
	if err != nil {
		log.Err(err).Msg("problem with deleting sample")
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, M{
		"message": "sample deleted",
	})
}
