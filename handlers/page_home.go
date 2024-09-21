package handlers

import (
	"collector/components/pages"
	"net/http"
)

func (s *router) HomePageHandler(w http.ResponseWriter, r *http.Request) {

	err := pages.Home().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
