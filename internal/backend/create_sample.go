package backend

import (
	"context"
	"io"

	"github.com/tunema-org/sound-function/model"
)

type CreateSampleParams struct {
	AuthorID       int
	Name           string
	BPM            int
	Key            model.SampleKey
	KeyScale       model.SampleKeyScale
	Time           int
	AudioFile      io.Reader
	AudioFileType  string
	Cover          io.Reader
	CoverImageType string
	Price          float64
	TagIDs         []int
}

func (b *Backend) CreateSample(ctx context.Context, accessToken string, params CreateSampleParams) (int, error) {
	return 0, nil
}
