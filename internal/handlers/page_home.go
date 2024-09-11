package handlers

import (
	"collector/internal/components"
	"net/http"
)

func (s *router) HomePageHandler(w http.ResponseWriter, r *http.Request) {

	err := components.HomePage().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
