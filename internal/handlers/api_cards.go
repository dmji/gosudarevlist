package handlers

import (
	"collector/components/cards"
	"collector/internal/services"
	"collector/pkg/custom_url"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type ApiCardsParams struct {
	Page        int
	SearchQuery string

	NextPage int
}

func NewParams(q url.Values, defaultPage int) ApiCardsParams {

	searchQueryStr := q.Get("query")
	pageStr := q.Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = defaultPage
		log.Print("Params ApiCards | Page argumen not passed")
	}

	return ApiCardsParams{
		Page:     page,
		NextPage: page + 1,

		SearchQuery: searchQueryStr,
	}
}

func (p *ApiCardsParams) ApplyToQuery(query *url.Values) {

	query.Set("page", strconv.Itoa(p.NextPage))

	if p.SearchQuery != "" {
		query.Set("query", p.SearchQuery)
	}
}

func (router *router) ApiCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	query := custom_url.QueryCustomParse(r.URL.Query())
	params := NewParams(query, 1)

	cardItems := router.s.GenerateCards(ctx,
		services.GenerateCardsOptions{
			Page:        params.Page,
			SearchQuery: params.SearchQuery,
		})

	log.Printf("Handler | ApiCards: page='%d' (len: %d)", params.Page, len(cardItems))
	params.ApplyToQuery(&query)
	nextPageParams := custom_url.QueryValuesToString(&query)
	log.Printf("Handler | ApiCards: nextQuery='%s'", nextPageParams)

	err := cards.ListItem(cardItems, r.URL.Path, nextPageParams).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
