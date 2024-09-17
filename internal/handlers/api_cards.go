package handlers

import (
	"collector/internal/components"
	"collector/internal/services"
	requestutils "collector/pkg/request_utils"
	"log"
	"net/http"
	"strconv"
)

func (router *router) ApiCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
		log.Print("Page argumen not passed")
	}

	cards := router.s.GenerateCards(ctx,
		services.GenerateCardsOptions{
			Page:        page,
			SearchQuery: "",
		},
	)

	log.Printf("Handler | ApiCards: page='%d' (len: %d)", page, len(cards))
	requestutils.LogQuery(r, "ApiCards")

	if len(cards) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = components.ListItem(cards, page+1).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
