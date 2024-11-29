package model

import (
	"context"
	"errors"
)

type Category string

func CategoryFromString(s string) (Category, error) {
	switch s {

	case string(Categories.Anime):
		return Categories.Anime, nil
	case string(Categories.AnimeHentai):
		return Categories.AnimeHentai, nil
	case string(Categories.Manga):
		return Categories.Manga, nil
	case string(Categories.MangaHentai):
		return Categories.MangaHentai, nil
	case string(Categories.Music):
		return Categories.Music, nil
	case string(Categories.Dorama):
		return Categories.Dorama, nil
	case string(Categories.All):
		return Categories.All, nil
	}

	return Categories.Anime, errors.New("string not match any of categories")
}

type categories struct {
	Anime       Category
	AnimeHentai Category
	Manga       Category
	MangaHentai Category
	Music       Category
	Dorama      Category
	All         Category
}

// Categories - object to emulate enum class
var Categories = categories{
	Anime:       "anime",
	AnimeHentai: "anime_hentai",
	Manga:       "manga",
	MangaHentai: "manga_henai",
	Music:       "music",
	Dorama:      "dorama",
	All:         "",
}

func (c *Category) Presentation(ctx context.Context) string {
	switch *c {
	case Categories.Anime:
		return "Anime"
	case Categories.AnimeHentai:
		return "Anime Hentai"
	case Categories.Manga:
		return "Manga"
	case Categories.MangaHentai:
		return "Manga Hentai"
	case Categories.Music:
		return "Music"
	case Categories.Dorama:
		return "Dorama"
	case Categories.All:
		return ""
	default:
		return ""
	}
}
