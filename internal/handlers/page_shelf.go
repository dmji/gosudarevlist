package handlers

import (
	"collector/internal/components"
	requestutils "collector/pkg/request_utils"
	"log"
	"net/http"
)

func (router *router) ShelfPageHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Context()

	params := NewParams(r, 0)

	requestutils.LogQuery(r, "ShelfPageHandler")
	log.Printf("Handler | ShelfPageHandler params: %s", params.ToString())

	err := components.ListingPage(params.ToString()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
