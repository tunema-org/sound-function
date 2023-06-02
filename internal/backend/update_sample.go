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
	AudioFileType string
	CoverFile     io.Reader
	CoverFileType string
	Price         float64
	TagIDs        []int
}

func (b *Backend) UpdateSample(ctx context.Context, accessToken string, params UpdateSampleParams) (int, error) {
	_, claims, err := jwt.Verify(accessToken, b.cfg.JWTSecretKey)
	if err != nil {
		return 0, err
	}

	UserID, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid claims")

	}

	SampleID, err := b.repo.UpdateSample(ctx, repository.UpdateSampleParams{
		Sample: model.Sample{
			UserID:   int(UserID),
			Name:     params.Name,
			BPM:      params.BPM,
			Key:      params.Key,
			KeyScale: params.KeyScale,
			Time:     params.Time,
			Price:    params.Price,
		},
		TagIDs: params.TagIDs,
	})

	return SampleID, err
}
