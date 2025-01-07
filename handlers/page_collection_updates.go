package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/pages"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func (router *router) CollectionUpdatesPageHandler(cat model.Category) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, err := custom_url.Decode(r.URL.RawQuery, model.WithApiCardsParamsSetPage(1))
		if err != nil {
			logger.Errorw(ctx, "CollectionUpdatesPageHandler decode query failed", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nextPageParams, err := custom_url.Encode(params)
		if err != nil {
			logger.Errorw(ctx, "CollectionUpdatesPageHandler encode query failed", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Infow(ctx, "Handler | ShelfPageHandler params", "params", nextPageParams)

		err = pages.CollectionUpdates(
			"/api/updates",
			custom_url.QueryOrEmpty(nextPageParams),
			params.SearchQuery,
		).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
