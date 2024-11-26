package handlers

import (
	"collector/components/pages"
	"collector/internal/query_cards"
	"collector/pkg/custom_url"
	"collector/pkg/logger"
	"net/http"
)

func (router *router) CollectionListingPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prm := query_cards.Parse(ctx, r.URL.Query(), 1)

	query := prm.Values(ctx)
	nextPageParams := custom_url.QueryValuesToString(&query)

	logger.Infow(ctx, "Handler | ShelfPageHandler", "params", nextPageParams)

	err := pages.CollectionListing(
		"/api/cards",
		"?"+prm.Values(ctx).Encode(),
		prm.SearchQuery,
	).Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
