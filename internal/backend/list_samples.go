package backend

import (
	"context"

	"github.com/tunema-org/sound-function/internal/repository"
)

func (b *Backend) ListSamples(ctx context.Context) ([]repository.ListSamplesResult, error) {
	result, err := b.repo.ListSamples(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
