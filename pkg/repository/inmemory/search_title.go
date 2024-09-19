package repository_inmemory

import (
	"collector/pkg/model"
	"context"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func (r *repository) SearchTitle(ctx context.Context, title string) ([]model.AnimeLayerItem, error) {

	res := make([]model.AnimeLayerItem, 0, 10)

	for _, t := range r.db {

		if fuzzy.Match(title, t.Name) {
			res = append(res, t)
		}

	}

	return res, nil
}
