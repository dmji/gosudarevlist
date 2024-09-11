package handlers

import (
	"collector/internal/components"
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

	cards := router.s.GenerateCards(ctx, page)
	log.Printf("Page: %d (len: %d)", page, len(cards))

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
