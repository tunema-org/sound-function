package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tunema-org/sound-function/internal/backend"
	"github.com/tunema-org/sound-function/internal/repository"
)

func (h *handler) SearchSamples(c *gin.Context) {
	sampleName, tags, orderBy := c.Query("name"), c.Query("tags"), c.Query("order_by")

	strSplitTags := strings.Split(tags, ",")

	var splitTags []int
	if tags != "" && len(strSplitTags) != 0 {
		for _, tag := range strSplitTags {
			tagInt, err := strconv.Atoi(tag)
			if err != nil {
				c.JSON(http.StatusBadRequest, M{
					"message": "invalid tag",
				})
				return
			}

			splitTags = append(splitTags, tagInt)
		}
	}

	samples, err := h.backend.SearchSamples(c.Request.Context(), backend.SearchSamplesParams{
		Name:    sampleName,
		Tags:    splitTags,
		OrderBy: orderBy,
	})
	if err != nil {
		if errors.Is(err, repository.ErrTagDoesNotExist) {
			c.JSON(http.StatusBadRequest, M{
				"message": "invalid tag",
			})
			return
		}

		log.Err(err).Msg("problem with searching samples")
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
