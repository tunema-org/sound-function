package backend

import (
	"context"
	"strings"

	"github.com/tunema-org/sound-function/internal/repository"
)

type SearchSamplesParams struct {
	Name    string
	Tags    []int
	OrderBy string
}

func (b *Backend) SearchSamples(ctx context.Context, params SearchSamplesParams) ([]repository.SearchSamplesResult, error) {
	for _, tag := range params.Tags {
		if _, err := b.repo.GetTag(ctx, tag); err != nil {
			return nil, repository.ErrTagDoesNotExist
		}
	}

	result, err := b.repo.SearchSamples(ctx, strings.ToLower(params.Name), params.Tags, params.OrderBy)
	if err != nil {
		return nil, err
	}

	return result, nil
}
