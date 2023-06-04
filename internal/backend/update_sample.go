package backend

import (
	"context"
	"errors"
	"io"

	"github.com/tunema-org/sound-function/internal/jwt"
	"github.com/tunema-org/sound-function/internal/repository"
	"github.com/tunema-org/sound-function/model"
)

type UpdateSampleParams struct {
	Name          string
	BPM           int
	Key           model.SampleKey
	KeyScale      model.SampleKeyScale
	Time          int
	CoverFile     io.Reader
	CoverFileType string
	Price         float64
	TagIDs        []int
}

func (b *Backend) UpdateSample(ctx context.Context, accessToken string, sampleID int, params UpdateSampleParams) (int, error) {
	_, claims, err := jwt.Verify(accessToken, b.cfg.JWTSecretKey)
	if err != nil {
		return 0, err
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	err = b.repo.UpdateSample(ctx, sampleID, repository.UpdateSampleParams{
		Sample: model.Sample{
			UserID:   int(userID),
			Name:     params.Name,
			BPM:      params.BPM,
			Key:      params.Key,
			KeyScale: params.KeyScale,
			Time:     params.Time,
			Price:    params.Price,
		},
		TagIDs: params.TagIDs,
	})
	if err != nil {
		return 0, err
	}

	return sampleID, nil
}
