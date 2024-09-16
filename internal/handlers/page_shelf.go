package handlers

import (
	"collector/internal/components"
	requestutils "collector/pkg/request_utils"
	"net/http"
)

func (router *router) ShelfPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	requestutils.LogQuery(r, "ShelfPageHandler")
	initialCards := router.s.GenerateCards(ctx, 1)

	err := components.ListingPage(initialCards, 2).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
