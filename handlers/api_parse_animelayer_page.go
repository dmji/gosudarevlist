package handlers

import (
	"collector/components/pages"
	"collector/internal/animelayer_parser"
	animelayer_service "collector/pkg/animelayer/service"
	"collector/pkg/custom_url"
	"collector/pkg/parser"
	"fmt"
	"log"
	"net/http"
)

func (router *router) ApiMyAnimeListParsePage(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	query := custom_url.QueryCustomParse(r.URL.Query())
	guid := query.Get("guid")

	/* 	cardItems := router.s.GenerateCards(ctx,
	services.GenerateCardsOptions{
		Page:        params.Page,
		SearchQuery: params.SearchQuery,
	})
	*/

	url := animelayer_service.FormatUrlToItem(guid)

	client, err := parser.HttpClientWithAuth(
		animelayer_service.BaseUrl,
		getTestCreadentials(),
	)

	if err != nil {
		panic(err)
	}

	doc, err := parser.LoadHtmlDocument(client, url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Recursive parsing
	item := animelayer_parser.ParseItem(ctx, doc)

	info := pages.ScanResultDescription{
		GUID: item.GUID,
	}
	info.Descriptions = append(info.Descriptions, "CreatedDate: "+item.CreatedDate)
	info.Descriptions = append(info.Descriptions, "UpdatedDate: "+item.UpdatedDate)
	info.Descriptions = append(info.Descriptions, "LastCheckedDate: "+item.LastCheckedDate)
	info.Descriptions = append(info.Descriptions, "RefImageCover: "+item.RefImageCover)
	info.Descriptions = append(info.Descriptions, "RefImagePreview: "+item.RefImagePreview)
	info.Descriptions = append(info.Descriptions, "TorrentFilesSize: "+item.TorrentFilesSize)
	for _, d := range item.Descriptions {
		info.Descriptions = append(info.Descriptions, d.Key+": "+d.Value)
	}
	log.Printf("Handler | ApiMyAnimeListParsePage: guid='%s'", guid)

	err = pages.ResultsDescription(info).Render(ctx, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
