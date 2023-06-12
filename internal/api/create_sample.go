package api

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tunema-org/sound-function/internal/backend"
	"github.com/tunema-org/sound-function/internal/mime"
	"github.com/tunema-org/sound-function/model"
	"golang.org/x/exp/slices"
)

type CreateSampleInput struct {
	Name     string  `form:"name" binding:"required"`
	BPM      int     `form:"bpm" binding:"required"`
	Key      string  `form:"key" binding:"required"`
	KeyScale string  `form:"key_scale" binding:"required"`
	Time     int     `form:"time" binding:"required"`
	Price    float64 `form:"price" binding:"required"`
	TagIDs   []int   `form:"tag_ids[]" binding:"required"`
}

var supportedSampleAudioFileTypes = []string{
	"audio/mpeg",
	"audio/wav",
	"audio/wave",
	"audio/aac",
	"audio/ogg",
	"application/octet-stream",
}

var supportedSampleCoverImageTypes = []string{
	"image/jpeg",
	"image/jpg",
	"image/png",
	"image/webp",
	"application/octet-stream",
}

func (i CreateSampleInput) Validate() error {
	if err := requiredFields(map[string]any{
		"name":      i.Name,
		"bpm":       i.BPM,
		"key":       i.Key,
		"key_scale": i.KeyScale,
		"time":      i.Time,
		"price":     i.Price,
		"tag_ids[]": i.TagIDs,
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

func (h *handler) CreateSample(c *gin.Context) {

	authorization := strings.Split(c.Request.Header["Authorization"][0], " ")
	if len(authorization) != 2 {
		c.JSON(http.StatusUnauthorized, M{
			"message": "please login",
		})
		return
	}

	var input CreateSampleInput

	if err := c.Bind(&input); err != nil {
		log.Err(err).Msg(err.Error())
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

	audioFileHeader, err := c.FormFile("audio_file")
	if err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(http.StatusUnprocessableEntity, M{
			"message": "invalid audio file",
		})
		return
	}

	coverFileHeader, err := c.FormFile("cover_file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, M{
			"message": "invalid cover file",
		})
		return
	}

	audioFile, err := audioFileHeader.Open()
	if err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(http.StatusUnprocessableEntity, M{
			"message": "invalid audio file",
		})
		return
	}

	coverFile, err := coverFileHeader.Open()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, M{
			"message": "invalid cover file",
		})
		return
	}

	// max cover file size 5mb
	if coverFileHeader.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, M{
			"message": "max cover image size is 5mb",
		})
		return
	}

	if !mime.Contains(audioFile, supportedSampleAudioFileTypes) {
		log.Info().Msg(audioFileHeader.Header.Get("Content-Type"))
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid audio file type",
		})
		return
	}

	if !mime.Contains(coverFile, supportedSampleCoverImageTypes) {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid cover file type",
		})
		return
	}

	sampleID, err := h.backend.CreateSample(c.Request.Context(), authorization[1], backend.CreateSampleParams{
		Name:          input.Name,
		BPM:           input.BPM,
		Key:           model.SampleKey(input.Key),
		KeyScale:      model.SampleKeyScale(input.KeyScale),
		Time:          input.Time,
		AudioFile:     audioFile,
		AudioFileType: filepath.Ext(audioFileHeader.Filename),
		CoverFile:     coverFile,
		CoverFileType: filepath.Ext(coverFileHeader.Filename),
		Price:         input.Price,
		TagIDs:        input.TagIDs,
	})
	if err != nil {
		log.Err(err).Msg("problem with creating sample")
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, M{
		"message":   "sample created",
		"sample_id": sampleID,
	})
}
