package backend

import (
	"context"
)

type ListTagsAndCategoriesOutputTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListTagsAndCategoriesOutput struct {
	Category string                           `json:"category"`
	Tags     []ListTagsAndCategoriesOutputTag `json:"tags"`
}

func (b *Backend) ListTagsAndCategories(ctx context.Context) ([]ListTagsAndCategoriesOutput, error) {
	result, err := b.repo.ListTagsAndCategories(ctx)
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[string][]ListTagsAndCategoriesOutputTag)
	var output []ListTagsAndCategoriesOutput

	for _, item := range result {
		var tag ListTagsAndCategoriesOutputTag

		tag.ID = item.TagID
		tag.Name = item.TagName

		if tags, ok := categoryMap[item.CategoryName]; ok {
			categoryMap[item.CategoryName] = append(tags, tag)
		} else {
			categoryMap[item.CategoryName] = []ListTagsAndCategoriesOutputTag{tag}
		}
	}

	for category, tags := range categoryMap {
		output = append(output, ListTagsAndCategoriesOutput{
			Category: category,
			Tags:     tags,
		})
	}

	return output, nil
}
