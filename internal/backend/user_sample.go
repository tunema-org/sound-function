package backend

import (
	"context"

	"github.com/tunema-org/sound-function/internal/repository"
)

func (b *Backend) UserSamples(ctx context.Context) ([]repository.UserSamples, error) {
	result, err := b.repo.UserSamples(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
