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

type ListTagsAndCategoriesResult struct {
	TagID        int    `json:"tag_id"`
	TagName      string `json:"tag_name"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func (r *Repository) ListTagsAndCategories(ctx context.Context) ([]ListTagsAndCategoriesResult, error) {
	var result []ListTagsAndCategoriesResult

	query := `
	SELECT
		tags.id,
		tags.name,
		categories.id,
		categories.name
	FROM
		tags
	LEFT JOIN categories ON tags.category_id = categories.id;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row ListTagsAndCategoriesResult

		err := rows.Scan(&row.TagID, &row.TagName, &row.CategoryID, &row.CategoryName)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}
