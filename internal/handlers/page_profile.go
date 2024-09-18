package handlers

import (
	"collector/components/pages"
	requestutils "collector/pkg/request_utils"
	"net/http"
)

func (s *router) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {

	requestutils.LogQuery(r, "ProfilePageHandler")
	err := pages.Profile().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
