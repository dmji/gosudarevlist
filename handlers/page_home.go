package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/pages"
)

func (s *router) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/animelayer/anime", http.StatusMovedPermanently)

	err := pages.Home().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
