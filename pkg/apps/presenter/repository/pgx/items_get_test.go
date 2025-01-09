package repository_pgx_test

import (
	"context"
	"testing"

	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
)

func TestGetITemsByCategory(t *testing.T) {
	repo, ctx := InitRepo(context.Background())

	items, _ := repo.GetItems(ctx, model.OptionsGetItems{
		PageIndex:       1,
		CountForOnePage: 20,

		SearchQuery:         "",
		SimilarityThreshold: 0.05,

		Categories: []model.Category{},
		Statuses:   []model.ReleaseStatus{},
	})
	println(items)
}

func TestGetFiltersByCategory(t *testing.T) {
	repo, ctx := InitRepo(context.Background())

	items, _ := repo.GetFilters(ctx, model.OptionsGetItems{
		PageIndex:       0,
		CountForOnePage: 20,

		SearchQuery:         "",
		SimilarityThreshold: 0.05,

		Categories: []model.Category{
			model.CategoryAnime,
			model.CategoryAnimeHentai,
		},
	})

	println(items)
}