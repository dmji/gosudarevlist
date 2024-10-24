package handlers

import (
	"collector/components/pages"
	"collector/internal/filters"
	"collector/pkg/custom_url"
	"log"
	"net/http"
	"slices"
	"strings"
)

func (router *router) ShelfPageHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Context()

	query := custom_url.QueryCustomParse(r.URL.Query())

	searchField := query.Get("query")

	sts := strings.Split(query.Get("st"), "-")
	filterParams := filters.FilterParameters{
		SearchField: searchField,
		Categories: []filters.FilterCategory{
			{
				ParameterTitle: "st",
				Values: []filters.FilterValue{
					{
						DisplayName:   "On Air",
						ParameterName: "air",
						Checked:       slices.ContainsFunc(sts, func(s string) bool { return s == "air" }),
					},
				},
			},
		},
	}

	nextPageParams := custom_url.QueryValuesToString(&query)

	log.Printf("Handler | ShelfPageHandler params: %s", nextPageParams)

	err := pages.Gallery(
		filterParams,
		"/api/cards",
		nextPageParams,
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
