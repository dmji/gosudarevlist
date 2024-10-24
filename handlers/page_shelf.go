package handlers

import (
	"collector/components/pages"
	"collector/internal/filters"
	"collector/pkg/custom_url"
	"log"
	"net/http"
)

func (router *router) ShelfPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prm := filters.ParseApiCardsParams(ctx, r.URL.Query(), 1)
	filterParams := filters.NewFiltersState(prm)

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
