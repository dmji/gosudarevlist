package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/pages"
	"github.com/dmji/gosudarevlist/internal/query_cards"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (router *router) CollectionUpdatesPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prm := query_cards.Parse(ctx, r.URL.Query(), 1)

	query := prm.Values(ctx)
	nextPageParams := custom_url.QueryValuesToString(&query)

	logger.Infow(ctx, "Handler | ShelfPageHandler params", "params", nextPageParams)

	err := pages.CollectionUpdates(
		"/api/updates",
		nextPageParams,
		prm.SearchQuery,
	).Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
