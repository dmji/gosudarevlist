package handlers

import (
	"net/http"
	"strconv"

	"github.com/dmji/gosudarevlist/components/pages"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *router) MalItemPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		logger.Errorw(ctx, "PathValue parsing failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := s.mal.GetItem(ctx, int(id))
	err = pages.MalAnimeCard(t).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
