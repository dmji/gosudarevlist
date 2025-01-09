package repository_pgx_test

import (
	"context"
	"testing"

	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

func TestGetITemsByCategory(t *testing.T) {
	repo, ctx := InitRepo(context.Background())

	items, _ := repo.GetItems(ctx, model.OptionsGetItems{
		PageIndex:       1,
		CountForOnePage: 20,

		SearchQuery:         "",
		SimilarityThreshold: 0.05,

		Categories: []enums.Category{},
		Statuses:   []enums.ReleaseStatus{},
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

		Categories: []enums.Category{
			enums.CategoryAnime,
			enums.CategoryAnimeHentai,
		},
	})

	println(items)
}
