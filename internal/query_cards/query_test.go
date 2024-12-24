package query_cards_test

import (
	"context"
	"net/url"
	"slices"
	"testing"

	"github.com/dmji/gosudarevlist/internal/query_cards"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func TestQueryWriteRead(t *testing.T) {
	opt := query_cards.ApiCardsParams{
		Page:        2,
		SearchQuery: "Worry",
		Categories: []model.Category{
			model.Categories.Anime,
			model.Categories.AnimeHentai,
		},
		Statuses: []model.Status{
			model.Statuses.OnAir,
		},
	}

	v := opt.Values(context.Background())
	s := v.Encode()

	u, _ := url.ParseQuery(s)
	qn := query_cards.Parse(context.Background(), u, 1)

	if opt.Page != qn.Page {
		t.Fatalf("expected %v, got %v", opt.Page, qn.Page)
	}
	if opt.SearchQuery != qn.SearchQuery {
		t.Fatalf("expected %v, got %v", opt.SearchQuery, qn.SearchQuery)
	}
	if !slices.Equal(opt.Categories, qn.Categories) {
		t.Fatalf("expected %v, got %v", opt.Categories, qn.Categories)
	}

	if !slices.Equal(opt.Statuses, qn.Statuses) {
		t.Fatalf("expected %v, got %v", opt.Categories, qn.Categories)
	}
}
