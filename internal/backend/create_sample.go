package backend

import (
	"context"
	"errors"
	"io"
	"path"

	"github.com/google/uuid"
	"github.com/tunema-org/sound-function/internal/jwt"
	"github.com/tunema-org/sound-function/internal/repository"
	"github.com/tunema-org/sound-function/model"
)

type CreateSampleParams struct {
	Name          string
	BPM           int
	Key           model.SampleKey
	KeyScale      model.SampleKeyScale
	Time          int
	AudioFile     io.Reader
	AudioFileType string
	CoverFile     io.Reader
	CoverFileType string
	Price         float64
	TagIDs        []int
}

func (b *Backend) CreateSample(ctx context.Context, accessToken string, params CreateSampleParams) (int, error) {
	storeSampleFileResult, err := b.storeSampleFiles(ctx, params)
	if err != nil {
		return 0, err
	}

	_, claims, err := jwt.Verify(accessToken, b.cfg.JWTSecretKey)
	if err != nil {
		return 0, err
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	sampleID, err := b.repo.InsertSample(ctx, repository.InsertSampleParams{
		Sample: model.Sample{
			UserID:   int(userID),
			Name:     params.Name,
			BPM:      params.BPM,
			Key:      params.Key,
			KeyScale: params.KeyScale,
			Time:     params.Time,
			Price:    params.Price,
			FileURL:  storeSampleFileResult.AudioFileURL,
			CoverURL: storeSampleFileResult.CoverFileURL,
		},
		TagIDs: params.TagIDs,
	})
	if err != nil {
		return 0, err
	}

	return sampleID, nil
}

type storeSampleFilesResult struct {
	AudioFileURL string
	AudioFileKey string
	CoverFileURL string
	CoverFileKey string
}

func (b *Backend) storeSampleFiles(ctx context.Context, params CreateSampleParams) (storeSampleFilesResult, error) {
	var result storeSampleFilesResult

	sampleFileKey := uuid.New().String()

	result.AudioFileKey = path.Join("samples", sampleFileKey+params.AudioFileType)
	result.CoverFileKey = path.Join("samples", sampleFileKey+params.CoverFileType)

	s3AudioFileUploadOutput, err := b.clients.S3.UploadFile(ctx, result.AudioFileKey, params.AudioFile)
	if err != nil {
		return storeSampleFilesResult{}, err
	}

	result.AudioFileURL = s3AudioFileUploadOutput.Location

	s3CoverFileUploadOutput, err := b.clients.S3.UploadFile(ctx, result.CoverFileKey, params.CoverFile)
	if err != nil {
		b.clients.S3.DeleteFile(ctx, result.AudioFileKey)
		return storeSampleFilesResult{}, err
	}

	result.CoverFileURL = s3CoverFileUploadOutput.Location

	return result, nil
}
