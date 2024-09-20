package handlers

import (
	"collector/components/cards"
	"collector/components/pages"
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

	cats := strings.Split(query.Get("cat"), "-")
	sts := strings.Split(query.Get("st"), "-")
	filterParams := cards.FilterParameters{
		SearchField: searchField,
		Categories: []cards.FilterCategory{
			{
				DisplayTitle:   "Category",
				ParameterTitle: "cat",
				Values: []cards.FilterValue{
					{
						DisplayName:   "Action",
						ParameterName: "action",
						Checked:       slices.ContainsFunc(cats, func(s string) bool { return s == "action" }),
					},
					{
						DisplayName:   "Fantastic",
						ParameterName: "fantastic",
						Checked:       slices.ContainsFunc(cats, func(s string) bool { return s == "fantastic" }),
					},
					{
						DisplayName:   "Historic",
						ParameterName: "historic",
						Checked:       slices.ContainsFunc(cats, func(s string) bool { return s == "historic" }),
					},
					{
						DisplayName:   "Isekai",
						ParameterName: "isekai",
						Checked:       slices.ContainsFunc(cats, func(s string) bool { return s == "isekai" }),
					},
				},
			},
			{
				DisplayTitle:   "Status",
				ParameterTitle: "st",
				Values: []cards.FilterValue{
					{
						DisplayName:   "On Air",
						ParameterName: "air",
						Checked:       slices.ContainsFunc(sts, func(s string) bool { return s == "air" }),
					},
					{
						DisplayName:   "Completed",
						ParameterName: "completed",
						Checked:       slices.ContainsFunc(sts, func(s string) bool { return s == "completed" }),
					},
					{
						DisplayName:   "Not Started",
						ParameterName: "not_started",
						Checked:       slices.ContainsFunc(sts, func(s string) bool { return s == "not_started" }),
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
