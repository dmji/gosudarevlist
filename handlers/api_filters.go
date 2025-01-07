package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/cards"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/expose_header_utils"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func (router *router) ApiFilters(cat model.Category) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params, err := custom_url.Decode(r.URL.RawQuery, model.WithApiCardsParamsSetPage(1))
		if err != nil {
			logger.Errorw(ctx, "ApiFilters | Decode query failed", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		urlPushed, err := expose_header_utils.HxPushUrl(ctx, w, r, func(q string) (string, error) { return custom_url.Encode(&params) })
		if err != nil {
			logger.Errorw(ctx, "ApiFilters | Parameters push to url failed", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Infow(ctx, "ApiFilters | Decode query", "params", params, "query", urlPushed.String())
		items := router.s.GetFilters(ctx, params, cat)

		err = cards.FilterFlagsPopulate(items).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
