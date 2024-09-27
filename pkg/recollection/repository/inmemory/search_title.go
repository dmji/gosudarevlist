package repository_inmemory

import (
	animelayer_model "collector/pkg/animelayer/model"
	"context"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func (r *repository) SearchTitle(ctx context.Context, title string) ([]animelayer_model.Item, error) {

	res := make([]animelayer_model.Item, 0, 10)

	for _, t := range r.db {

		if fuzzy.Match(title, t.Title) {
			res = append(res, t)
		}

	}

	return res, nil
}
