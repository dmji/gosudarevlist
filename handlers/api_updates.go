package handlers

import (
	"collector/components/pages"
	"collector/internal/query_cards"
	"net/http"
)

func (s *router) UpdatesListHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	params := query_cards.Parse(ctx, r.URL.Query(), 1)

	items := s.s.GetUpdates(ctx, params)

	notes := make([]pages.Item, 0, len(items))
	for _, item := range items {
		notes = append(notes, pages.Item{
			Name:       item.Title,
			Identifier: item.Identifier,
			Status:     item.Status,
			Date:       item.Date,
			/* 			Changes: []pages.Change{
				{
					TextOld: item.ValueOld,
					TextNew: item.ValueNew,
				},
			}, */
		})
	}

	err := pages.FormList(notes).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
