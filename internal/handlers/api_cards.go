package handlers

import (
	"collector/internal/components"
	"collector/internal/services"
	requestutils "collector/pkg/request_utils"
	"log"
	"net/http"
	"strconv"
)

type ApiCardsParams struct {
	Page        int
	SearchQuery string

	NextPage int
}

func NewParams(r *http.Request, defaultPage int) ApiCardsParams {

	searchQueryStr := r.URL.Query().Get("query")
	pageStr := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = defaultPage
		log.Print("Page argumen not passed")
	}

	return ApiCardsParams{
		Page:     page,
		NextPage: page + 1,

		SearchQuery: searchQueryStr,
	}
}

func (p *ApiCardsParams) ToString() string {
	res := ""

	res += "page=" + strconv.Itoa(p.NextPage)
	if p.SearchQuery != "" {
		res += "&query=" + p.SearchQuery
	}

	return res
}

func (router *router) ApiCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := NewParams(r, 1)

	cards := router.s.GenerateCards(ctx,
		services.GenerateCardsOptions{
			Page:        params.Page,
			SearchQuery: params.SearchQuery,
		})

	log.Printf("Handler | ApiCards: page='%d' (len: %d)", params.Page, len(cards))
	requestutils.LogQuery(r, "ApiCards")
	log.Printf("Handler | ApiCards params: %s", params.ToString())

	if len(cards) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err := components.ListItem(cards, params.ToString()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
