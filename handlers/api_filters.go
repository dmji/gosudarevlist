package handlers

import (
	"collector/components/cards"
	"collector/internal/query_cards"
	"net/http"
)

func (router *router) ApiFilters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := query_cards.Parse(ctx, r.URL.Query(), 1)

	items := router.s.GetFilters(ctx, params)

	err := cards.FilterFlagsPopulate(items).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
