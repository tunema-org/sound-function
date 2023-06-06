package repository

import (
	"context"

	"github.com/tunema-org/sound-function/model"
)

type DeleteSampleParams struct {
	Sample model.Sample
}

func (r *Repository) DeleteSample(ctx context.Context, sampleID int) error {
	deleteSampleQuery := `DELETE FROM samples WHERE id = $1`

	_, err := r.db.Exec(ctx, deleteSampleQuery, sampleID)
	if err != nil {
		return err
	}

	return nil

}
