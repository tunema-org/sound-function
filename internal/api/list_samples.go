package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryInput struct {
	Tag_ID int
}

// error kykny
func (h *handler) ListSamples(c *gin.Context) {
	var input CategoryInput

	Tag_ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, M{
			"message": "invalid sample id",
		})
		return
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
		return
	}

	c.JSON(http.StatusOK, M{
		"message":  "sample retrieved",
		"sampleID": Tag_ID,
	})

}
