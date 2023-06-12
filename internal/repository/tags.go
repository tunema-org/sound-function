package repository

import (
	"context"

	"github.com/tunema-org/sound-function/model"
)

func (r *Repository) GetTag(ctx context.Context, tagID int) (model.Tag, error) {
	var tag model.Tag

	query := `SELECT id, name FROM tags WHERE id = $1;`

	err := r.db.QueryRow(ctx, query, tagID).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return model.Tag{}, err
	}

	return tag, nil
}

func (r *Repository) ListTags(ctx context.Context) ([]model.Tag, error) {
	var tags []model.Tag

	query := `SELECT id, name FROM tags;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag model.Tag

		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
