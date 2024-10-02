package repository_inmemory

import (
	"context"

	"github.com/dmji/go-animelayer-parser"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

func (r *repository) SearchTitle(ctx context.Context, title string) ([]animelayer.ItemPartial, error) {

	res := make([]animelayer.ItemPartial, 0, 10)

	for _, t := range r.db {

		if fuzzy.Match(title, t.Title) {
			res = append(res, t)
		}

	}

	return res, nil
}
