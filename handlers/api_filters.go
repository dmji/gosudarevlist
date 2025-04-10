package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/cards"
	"github.com/dmji/gosudarevlist/internal/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/expose_header_utils"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (router *router) ApiFilters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cat, err := enums.CategoryFromString(r.PathValue("category"))
	if err != nil {
		logger.Errorw(ctx, "PathValue parsing failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params, err := custom_url.Decode(r.URL.RawQuery, model.WithApiCardsParamsSetPage(1))
	if err != nil {
		logger.Errorw(ctx, "ApiFilters | Decode query failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = expose_header_utils.HxReplaceUrl(ctx, w, r, func(q string) (string, error) { return custom_url.Encode(&params) })
	if err != nil {
		logger.Errorw(ctx, "ApiFilters | Parameters push to url failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := router.presentService.GetFilters(ctx, params, cat)

	err = cards.FilterFlagsPopulate(items).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
