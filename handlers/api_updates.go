package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/cards"
	"github.com/dmji/gosudarevlist/internal/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *router) ApiUpdates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cat, err := enums.CategoryFromString(r.PathValue("category"))
	if err != nil {
		logger.Errorw(ctx, "PathValue parsing failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params, err := custom_url.Decode(r.URL.RawQuery, model.WithApiCardsParamsSetPage(1))
	if err != nil {
		logger.Errorw(ctx, "ApiUpdates decode query failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := s.presentService.GetUpdates(ctx, params, cat)

	params.Page += 1
	nextPageParams, err := custom_url.Encode(params)
	if err != nil {
		logger.Errorw(ctx, "ApiUpdates encode query failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = cards.CollectionUpdatesBatch(
		items,
		r.URL.Path,
		custom_url.QueryOrEmpty(nextPageParams),
		true,
		int(params.Page)-1,
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
