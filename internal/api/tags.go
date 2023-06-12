package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListTags(c *gin.Context) {}

func (h *handler) ListTagsAndCategories(c *gin.Context) {
	result, err := h.backend.ListTagsAndCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, M{
		"message":     "success",
		"items":       result,
		"total_items": len(result),
	})
}
