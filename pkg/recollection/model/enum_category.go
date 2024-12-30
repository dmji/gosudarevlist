package model

import (
	"context"
	"errors"
)

type Category string

func (c *Category) String() string {
	return string(*c)
}

var Categories = struct {
	Anime       Category
	AnimeHentai Category
	Manga       Category
	MangaHentai Category
	Music       Category
	Dorama      Category
	All         Category
}{
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

func CategoryFromString(s string) (Category, error) {
	switch s {

	case Categories.Anime.String():
		return Categories.Anime, nil
	case Categories.AnimeHentai.String():
		return Categories.AnimeHentai, nil
	case Categories.Manga.String():
		return Categories.Manga, nil
	case Categories.MangaHentai.String():
		return Categories.MangaHentai, nil
	case Categories.Music.String():
		return Categories.Music, nil
	case Categories.Dorama.String():
		return Categories.Dorama, nil
	case Categories.All.String():
		return Categories.All, nil
	}

	return Categories.Anime, errors.New("string not match any of categories")
}
