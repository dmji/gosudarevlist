package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/pages"
	"github.com/dmji/gosudarevlist/internal/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (router *router) CollectionUpdatesPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cat, err := enums.CategoryFromString(r.PathValue("category"))
	if err != nil {
		logger.Errorw(ctx, "PathValue parsing failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	categoriesLine := make([]model.CategoryButton, 0, len(categories))
	for _, c := range categories {
		categoriesLine = append(categoriesLine, model.NewCategoryButton(c,
			"/animelayer/%s",
			"/animelayer/%s/updates",
			c == cat,
		))
	}

	err = pages.CollectionUpdates(
		"/api/filters/"+cat.String(),
		"/api/updates/"+cat.String(),
		custom_url.QueryOrEmpty(nextPageParams),
		categoriesLine,
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
