package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/tunema-org/sound-function/model"
)

type UpdateSampleParams struct {
	Sample model.Sample
}

func (r *Repository) UpdateSample(ctx context.Context, params UpdateSampleParams) (int, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}

	updateSampleQuery := `UPDATE samples SET name = $1, bpm = $2, key = $4, key_scale = $5, time = $6, price = $7 WHERE id = $8`

	var SampleID int
	err = tx.QueryRow(ctx, updateSampleQuery,
		params.Sample.Name,
		params.Sample.BPM,
		params.Sample.Key,
		params.Sample.KeyScale,
		params.Sample.Time,
		params.Sample.Price,
	).Scan(&SampleID)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	fmt.Println("Sample Updated")

	return 0, err
}
