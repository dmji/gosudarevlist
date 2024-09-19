package handlers

import (
	"collector/components/pages"
	"log"
	"net/http"
)

func (router *router) ShelfPageHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Context()

	params := NewParams(r, 0)

	log.Printf("Handler | ShelfPageHandler params: %s", params.ToString())

	err := pages.Listing(params.ToString()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
