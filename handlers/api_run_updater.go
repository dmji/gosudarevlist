package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *router) RunUpdaterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cat, err := enums.CategoryFromString(r.PathValue("category"))
	if err != nil {
		logger.Errorw(ctx, "PathValue parsing failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Infow(r.Context(), "RunUpdaterHandler recived", "category", cat, "rurl", r.RequestURI, "url", r.URL, "header", r.Header)
}
