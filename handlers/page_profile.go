package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/pages"
)

func (s *router) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {

	err := pages.Profile(r.Context()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
