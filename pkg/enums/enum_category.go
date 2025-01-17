package enums

//go:generate go run github.com/dmji/go-stringer@latest -type=Category -trimprefix=Category -output enum_category_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

import (
	"context"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Category int8

const (
	CategoryAnime Category = iota
	CategoryAnimeHentai
	CategoryManga
	CategoryMangaHentai
	CategoryMusic
	CategoryDorama
	CategoryAll
)

func (c Category) Presentation(ctx context.Context) string {
	switch c {
	case CategoryAnime:
		return lang.Message(ctx, &i18n.Message{
			ID:    _CategoryAnime_i18n_ID,
			Other: "Anime",
		})
	case CategoryAnimeHentai:
		return lang.Message(ctx, &i18n.Message{
			ID:    _CategoryAnimeHentai_i18n_ID,
			Other: "Anime Hentai",
		})
	case CategoryManga:
		return lang.Message(ctx, &i18n.Message{
			ID:    _CategoryManga_i18n_ID,
			Other: "Manga",
		})
	case CategoryMangaHentai:
		return lang.Message(ctx, &i18n.Message{
			ID:    _CategoryMangaHentai_i18n_ID,
			Other: "Manga Hentai",
		})
	case CategoryMusic:
		return lang.Message(ctx, &i18n.Message{
			ID:    _CategoryMusic_i18n_ID,
			Other: "Music",
		})
	case CategoryDorama:
		return lang.Message(ctx, &i18n.Message{
			ID:    _CategoryDorama_i18n_ID,
			Other: "Dorama",
		})
	case CategoryAll:
		return lang.Message(ctx, &i18n.Message{
			ID:    _CategoryAll_i18n_ID,
			Other: "All",
		})
	default:
		return ""
	}
}
