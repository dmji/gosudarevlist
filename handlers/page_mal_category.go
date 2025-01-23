package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/pages"
)

func (s *router) MalCategoryPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cat := r.PathValue("category")
	_ = cat

	t, err := s.mal.GetCategory(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pages.MalAnimeList(t).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
