package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tunema-org/sound-function/model"
	"golang.org/x/exp/slices"
)

type UpdateSampleInput struct {
	Name     string  `form:"name" binding:"required"`
	BPM      int     `form:"bpm" binding:"required"`
	Key      string  `form:"key" binding:"required"`
	KeyScale string  `form:"key_scale" binding:"required"`
	Time     int     `form:"time" binding:"required"`
	Price    float64 `form:"price" binding:"required"`
}

func (i UpdateSampleInput) Validate() error {
	if err := requiredFields(map[string]any{
		"name":      i.Name,
		"bpm":       i.BPM,
		"key":       i.Key,
		"key_scale": i.KeyScale,
		"time":      i.Time,
		"price":     i.Price,
	}); err != nil {
		return err
	}

	if !slices.Contains(model.SampleValidKeys, model.SampleKey(i.Key)) {
		return fmt.Errorf("key is invalid")
	}

	if !slices.Contains(model.SampleValidKeyScales, model.SampleKeyScale(i.KeyScale)) {
		return fmt.Errorf("key_scale is invalid")
	}

	return nil
}

func (h *handler) UpdateSample(c *gin.Context) {

	sampleID, err := strconv.Atoi(c.Param("id"))
	var input UpdateSampleInput

	if err != nil {
		c.JSON(400, M{
			"message": "invalid sample id",
		})
		return
	}

	if err != nil {
		log.Err(err).Msg("problem with updating sample")
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
		return
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, M{
			"message": err.Error(),
		})
		return
	}

	authorization := strings.Split(c.Request.Header["Authorization"][0], " ")
	if len(authorization) != 2 {
		c.JSON(http.StatusUnauthorized, M{
			"message": "please login",
		})
		return
	}

	c.JSON(http.StatusCreated, M{
		"message":  "sample updated",
		"sampleID": sampleID,
	})
}
