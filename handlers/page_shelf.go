package handlers

import (
	"collector/components/pages"
	"collector/internal/filter_cards"
	"collector/internal/query_cards"
	"collector/pkg/custom_url"
	"log"
	"net/http"
)

func (router *router) ShelfPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prm := query_cards.Parse(ctx, r.URL.Query(), 1)
	filterParams := filter_cards.NewFiltersState(prm)

	query := prm.Values(ctx)
	nextPageParams := custom_url.QueryValuesToString(&query)

	log.Printf("Handler | ShelfPageHandler params: %s", nextPageParams)

	err := pages.Gallery(
		filterParams,
		"/api/cards",
		nextPageParams,
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
