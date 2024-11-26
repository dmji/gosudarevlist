package query_cards_test

import (
	"collector/internal/query_cards"
	"collector/pkg/recollection/model"
	"context"
	"net/url"
	"slices"
	"testing"
)

func TestQueryWriteRead(t *testing.T) {
	opt := query_cards.ApiCardsParams{
		Page:        2,
		SearchQuery: "Worry",
		/* 		IsCompleted: &custom_types.BoolExProp{
			Value: custom_types.BoolExFalse,
			Name:  "completed",
		}, */
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

	/* if *opt.IsCompleted != *qn.IsCompleted {
		t.Fatalf("expected %v, got %v", opt.IsCompleted, qn.IsCompleted)
	}
	*/
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
