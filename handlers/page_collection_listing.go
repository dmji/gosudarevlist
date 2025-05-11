package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/pages"
	"github.com/dmji/gosudarevlist/internal/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

//var categories = []enums.Category{
//	enums.CategoryAnime,
//	enums.CategoryManga,
//	enums.CategoryDorama,
//	enums.CategoryMusic,
//}

func (router *router) CollectionListingPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cat, err := enums.CategoryFromString(r.PathValue("category"))
	if err != nil {
		logger.Errorw(ctx, "PathValue parsing failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params, err := custom_url.Decode(r.URL.RawQuery, model.WithApiCardsParamsSetPage(1))
	if err != nil {
		logger.Errorw(ctx, "CollectionListingPageHandler decode query failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nextPageParams, err := custom_url.Encode(params)
	if err != nil {
		logger.Errorw(ctx, "CollectionListingPageHandler encode query failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Infow(ctx, "Handler | ShelfPageHandler", "params", nextPageParams)

	//categoriesLine := make([]model.CategoryButton, 0, len(categories))
	//for _, c := range categories {
	//	categoriesLine = append(categoriesLine, model.NewCategoryButton(c,
	//		"/animelayer/%s",
	//		"/animelayer/%s/updates",
	//		c == cat,
	//	))
	//}

	err = pages.CollectionListing(
		"/api/filters/"+cat.String(),
		"/api/cards/"+cat.String(),
		nil, // categoriesLine
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
