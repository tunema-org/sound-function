package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tunema-org/sound-function/model"
	"golang.org/x/exp/slices"
)

func (h *handler) UpdateSample(c *gin.Context) {}

type UpdateSampleInput struct {
	Name     string  `form:"name" binding:"required"`
	BPM      int     `form:"bpm" binding:"required"`
	Key      string  `form:"key" binding:"required"`
	KeyScale string  `form:"key_scale" binding:"required"`
	Time     int     `form:"time" binding:"required"`
	Price    float64 `form:"price" binding:"required"`
}

//sy bingung mau reuse supported cover img ny bgmn

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
