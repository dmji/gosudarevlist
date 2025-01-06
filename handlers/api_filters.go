package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/cards"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func (router *router) ApiFilters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params, err := custom_url.Decode(r.URL.RawQuery, model.WithApiCardsParamsSetPage(1))
	if err != nil {
		logger.Errorw(ctx, "ApiFilters decode query failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Infow(ctx, "ApiFilters decode query", "query", params)
	items := router.s.GetFilters(ctx, params)

	err = cards.FilterFlagsPopulate(items).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
