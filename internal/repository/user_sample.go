package repository

import (
	"context"

	"github.com/tunema-org/sound-function/model"
)

type UserSamples struct {
	Data       model.Sample `json:"data"`
	Tags       []string     `json:"tags"`
	ArtistName string       `json:"artist_name"`
	Sold       int          `json:"sold"`
}

func (r *Repository) UserSamples(ctx context.Context) ([]UserSamples, error) {
	var result []UserSamples

	query := `SELECT samples.*
	FROM
		samples
		LEFT JOIN users ON samples.user_id = users.id
		LEFT JOIN sample_tags ON samples.id = sample_tags.sample_id
		LEFT JOIN tags ON sample_tags.tag_id = tags.id
		LEFT JOIN order_products ON order_products.sample_id = samples.id
	WHERE
		users.id = $1
	GROUP BY
		samples.id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row UserSamples

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
