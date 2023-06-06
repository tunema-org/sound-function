package backend

import "context"

func (b *Backend) DeleteSample(ctx context.Context, sampleID int) error {
	err := b.repo.DeleteSample(ctx, sampleID)
	if err != nil {
		return err
	}

	return nil
}
