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

	var output []ListTagsAndCategoriesOutput

	for _, item := range result {
		var tag ListTagsAndCategoriesOutputTag

		tag.ID = item.TagID
		tag.Name = item.TagName

		var category ListTagsAndCategoriesOutput

		category.Category = item.CategoryName
		category.Tags = append(category.Tags, tag)

		output = append(output, category)
	}

	// TODO: OPTIMIZE THIS BULLSHIT
	for i := 0; i < len(output); i++ {
		for j := i + 1; j < len(output); j++ {
			if output[i].Category == output[j].Category {
				output[i].Tags = append(output[i].Tags, output[j].Tags...)
				output = append(output[:j], output[j+1:]...)
				j--
			}
		}
	}

	return output, nil
}
