package model

//go:generate go-stringer -type=Category -trimprefix=Category -output enum_category_string.go -nametransform=snake_case_lower -fromstringgenfn -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

import (
	"context"
)

type Category int8

const (
	CategoryAnime Category = iota
	CategoryAnimeHentai
	CategoryManga
	CategoryMangaHentai
	CategoryMusic
	CategoryDorama
)

func (c *Category) Presentation(ctx context.Context) string {
	switch *c {
	case CategoryAnime:
		return "Anime"
	case CategoryAnimeHentai:
		return "Anime Hentai"
	case CategoryManga:
		return "Manga"
	case CategoryMangaHentai:
		return "Manga Hentai"
	case CategoryMusic:
		return "Music"
	case CategoryDorama:
		return "Dorama"
	default:
		return ""
	}
}
