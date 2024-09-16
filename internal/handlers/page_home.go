package handlers

import (
	"collector/internal/components"
	requestutils "collector/pkg/request_utils"
	"net/http"
)

func (s *router) HomePageHandler(w http.ResponseWriter, r *http.Request) {

	requestutils.LogQuery(r, "HomePageHandler")
	err := components.HomePage().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
