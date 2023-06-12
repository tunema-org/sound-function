package repository

import (
	"context"

	"github.com/tunema-org/sound-function/model"
)

type ListSamplesResult struct {
	Data       model.Sample `json:"data"`
	Tags       []string     `json:"tags"`
	ArtistName string       `json:"artist_name"`
	Sold       int          `json:"sold"`
}

func (r *Repository) ListSamples(ctx context.Context) ([]ListSamplesResult, error) {
	var result []ListSamplesResult

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
	GROUP BY
		samples.id,
		users.username
	ORDER BY samples.created_at DESC;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row ListSamplesResult

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
