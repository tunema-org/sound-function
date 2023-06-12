package backend

import (
	"context"

	"github.com/tunema-org/sound-function/internal/repository"
)

func (b *Backend) ListUserSamples(ctx context.Context, userID int) ([]repository.UserSamples, error) {
	result, err := b.repo.ListUserSamples(ctx, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
