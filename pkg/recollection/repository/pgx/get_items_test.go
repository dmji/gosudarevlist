package repository_pgx_test

import (
	"collector/pkg/recollection/model"
	"context"
	"testing"
)

func TestGetITemsByCategory(t *testing.T) {

	repo, ctx := InitRepo(context.Background())

	repo.GetItems(ctx, model.OptionsGetItems{
		PageIndex:       0,
		CountForOnePage: 20,
	})
}

func TestGetFiltersByCategory(t *testing.T) {

	repo, ctx := InitRepo(context.Background())

	items, _ := repo.GetFilters(ctx, model.OptionsGetItems{
		PageIndex:       0,
		CountForOnePage: 20,

		SearchQuery: "",
		Categories: []model.Category{
			model.Categories.Anime,
			model.Categories.AnimeHentai,
		},
	})

	println(items)
}
