package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tunema-org/sound-function/model"
)

type UpdateSampleParams struct {
	Sample model.Sample
	TagIDs []int
}

func (r *Repository) UpdateSample(ctx context.Context, sampleID int, params UpdateSampleParams) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	updateSampleQuery := `UPDATE samples SET name = $1, bpm = $2, key = $4, key_scale = $5, time = $6, price = $7 WHERE id = $8`

	_, err = tx.Exec(ctx, updateSampleQuery,
		params.Sample.Name,
		params.Sample.BPM,
		params.Sample.Key,
		params.Sample.KeyScale,
		params.Sample.Time,
		params.Sample.Price,
	)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
