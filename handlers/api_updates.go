package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/cards"
	"github.com/dmji/gosudarevlist/internal/query_cards"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
)

func (s *router) ApiUpdates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := query_cards.Parse(ctx, r.URL.Query(), 1)

	items := s.s.GetUpdates(ctx, params)

	params.Page += 1
	query := params.Values(ctx)
	nextPageParams := custom_url.QueryValuesToString(&query)

	err := cards.CollectionUpdatesBatch(ctx, items, r.URL.Path, nextPageParams, true, int(params.Page)-1).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
