package handlers

import (
	"collector/components/cards"
	"collector/internal/query_cards"
	"collector/pkg/custom_url"
	"net/http"
)

func (router *router) ApiCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := query_cards.Parse(ctx, r.URL.Query(), 1)

	cardItems := router.s.GetItems(ctx, params)

	params.Page += 1
	query := params.Values(ctx)
	nextPageParams := custom_url.QueryValuesToString(&query)

	err := cards.CollectionListingBatch(cardItems, r.URL.Path, nextPageParams, true, int(params.Page)-1).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
