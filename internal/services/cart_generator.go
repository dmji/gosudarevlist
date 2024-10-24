package services

import (
	"collector/internal/filters"
	"collector/pkg/recollection/model"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/dmji/go-animelayer-parser"
)

const (
	perPage = 20
)

func queryEncodeForMyAnimeList(name string) string {
	params := url.Values{}
	params.Add("q", name)
	params.Add("cat", "all")

	return params.Encode()
}

func queryPosterFromItem(description *animelayer.Item) string {

	if img := description.RefImageCover; len(img) > 0 {
		return img
	}

	return "/assets/no_image.jpg"
}

func monthToRussian(month time.Month) string {
	switch month {
	case time.January:
		return "Января"
	case time.February:
		return "Февраля"
	case time.March:
		return "Марта"
	case time.April:
		return "Апреля"
	case time.May:
		return "Мая"
	case time.June:
		return "Июня"
	case time.July:
		return "Июля"
	case time.August:
		return "Августа"
	case time.September:
		return "Сентября"
	case time.October:
		return "Октября"
	case time.November:
		return "Ноября"
	case time.December:
		return "Декабря"
	}
	return ""
}

func timePtrToString(t *time.Time) string {
	if t == nil {
		return ""
	}

	if t.Year() == time.Now().Year() {
		return fmt.Sprintf(t.Format("02 %s в 15:04"), monthToRussian(t.Month()))
	}

	return fmt.Sprintf(t.Format("02 %s 2006 в 15:04"), monthToRussian(t.Month()))
}

func (s *services) GenerateCards(ctx context.Context, opt filters.ApiCardsParams) []model.ItemCartData {
	startID := (opt.Page - 1) * perPage

	items, err := s.AnimeLayerRepositoryDriver.GetItems(ctx, model.OptionsGetItems{
		Count:       perPage,
		Offset:      int64(startID),
		SearchQuery: opt.SearchQuery,
		Category:    model.Categories.Anime,
	})

	if err != nil {
		return nil
	}

	cardItems := make([]model.ItemCartData, 0, perPage)
	for _, item := range items {
		cardItems = append(cardItems, model.ItemCartData{
			Title:         item.Title,
			Image:         queryPosterFromItem(&item),
			Description:   item.Notes,
			CreatedDate:   timePtrToString(item.Updated.CreatedDate),
			UpdatedDate:   timePtrToString(item.Updated.UpdatedDate),
			TorrentWeight: item.Metrics.FilesSize,
			AnimeLayerRef: fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.Identifier),
			IsCompleted:   item.IsCompleted,
		})
	}
	return cardItems
}
