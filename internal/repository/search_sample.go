package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/tunema-org/sound-function/model"
)

type SearchSamplesResult struct {
	Data       model.Sample `json:"data"`
	Tags       []string     `json:"tags"`
	ArtistName string       `json:"artist_name"`
	Sold       int          `json:"sold"`
}

func (r *Repository) SearchSamples(ctx context.Context, likeName string, tags []int, orderBy string) ([]SearchSamplesResult, error) {
	var result []SearchSamplesResult

	query := r.buildSearchSampleQuery(likeName, tags, orderBy)

	var rows pgx.Rows
	var err error

	var args []any

	if likeName != "" {
		args = append(args, likeName)
	}

	if len(tags) > 0 {
		args = append(args, tags)
	}

	if orderBy != "" {
		args = append(args, orderBy)
	}

	if len(args) > 0 {
		rows, err = r.db.Query(ctx, query, args...)
	} else {
		rows, err = r.db.Query(ctx, query)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row SearchSamplesResult

		err := rows.Scan(
			&row.Data.ID,
			&row.Data.UserID,
			&row.Data.Name,
			&row.Data.BPM,
			&row.Data.Key,
			&row.Data.KeyScale,
			&row.Data.Time,
			&row.Data.FileURL,
			&row.Data.CoverURL,
			&row.Data.Price,
			&row.Data.CreatedAt,
			&row.ArtistName,
			&row.Tags,
			&row.Sold,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}

func (r *Repository) buildSearchSampleQuery(likeName string, tags []int, orderBy string) string {
	placeholderNo := 1

	query := `SELECT
		samples.*,
		users.username AS artist_name,
		ARRAY_AGG(tags.name) AS tags,
		COUNT(order_products.sample_id) AS sold
	FROM
		samples
		LEFT JOIN users ON samples.user_id = users.id
		LEFT JOIN sample_tags ON samples.id = sample_tags.sample_id
		LEFT JOIN tags ON sample_tags.tag_id = tags.id
		LEFT JOIN order_products ON order_products.sample_id = samples.id
		`

	if likeName != "" {
		query += fmt.Sprintf(`WHERE samples.name ILIKE '%%' || $%d || '%%'`, placeholderNo)
		placeholderNo++
	}

	query += `
	GROUP BY
		samples.id,
		users.username
	`

	if len(tags) > 0 {
		query += fmt.Sprintf("HAVING ARRAY_AGG(tags.id) <@ $%d", placeholderNo)
		placeholderNo++
	}

	if orderBy != "" {
		query += fmt.Sprintf("ORDER BY $%d DESC;", placeholderNo)
	} else {
		query += "ORDER BY samples.created_at DESC;"
	}

	return query
}
