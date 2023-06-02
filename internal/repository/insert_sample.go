package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tunema-org/sound-function/model"
)

type InsertSampleParams struct {
	Sample   model.Sample
	TagIDs   []int
	AuthorID int
}

func (r *Repository) InsertSample(ctx context.Context, params InsertSampleParams) (int, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}

	insertSampleQuery := `INSERT INTO samples(user_id, name, bpm, key, key_scale, time, file_url, cover_url, price, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW()) RETURNING id`

	var sampleID int
	err = tx.QueryRow(ctx, insertSampleQuery,
		params.AuthorID,
		params.Sample.Name,
		params.Sample.BPM,
		params.Sample.BPM,
	).Scan(&sampleID)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	insertSampleTagsQuery := `INSERT INTO sample_tags(sample_id, tag_id) VALUES ($1, $2)`
	var insertSampleTagErr error
	for _, tagID := range params.TagIDs {
		_, insertSampleTagErr = tx.Exec(ctx, insertSampleTagsQuery, sampleID, tagID)
		if insertSampleTagErr != nil {
			tx.Rollback(ctx)
			break
		}
	}

	if insertSampleTagErr != nil {
		return 0, insertSampleTagErr
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return sampleID, nil
}
