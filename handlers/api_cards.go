package handlers

import (
	"collector/components/cards"
	"collector/internal/filters"
	"collector/pkg/custom_url"
	"log"
	"net/http"
)

func (router *router) ApiCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := filters.ParseApiCardsParams(ctx, r.URL.Query(), 1)

	cardItems := router.s.GenerateCards(ctx, params)

	log.Printf("Handler | ApiCards: page='%d' (len: %d)", params.Page, len(cardItems))

	params.Page += 1
	query := params.Values(ctx)
	nextPageParams := custom_url.QueryValuesToString(&query)

	log.Printf("Handler | ApiCards: nextQuery='%s'", nextPageParams)

	err := cards.ListItem(cardItems, r.URL.Path, nextPageParams, true).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
