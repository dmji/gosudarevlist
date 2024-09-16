package handlers

import (
	"collector/internal/components"
	requestutils "collector/pkg/request_utils"
	"net/http"
)

func (s *router) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {

	requestutils.LogQuery(r, "ProfilePageHandler")
	err := components.ProfilePage().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
