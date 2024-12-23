package env

import (
	"github.com/dmji/gosudarevlist/pkg/recollection/model"

	"github.com/dmji/go-animelayer-parser"
	"github.com/joho/godotenv"
)

func LoadEnv(deep int, canPanic bool) {
	path := ".env"
	for i := range deep {
		if i != 0 {
			path = "../" + path
		}
		err := godotenv.Load(path)
		if err == nil {
			return
		}
	}

	if canPanic {
		panic(".env not found")
	}
}

func StrToCategory(str string) animelayer.Category {
	switch str {
	case "anime":
		return animelayer.Categories.Anime()
	case "anime_hentai":
		return animelayer.Categories.AnimeHentai()
	case "manga":
		return animelayer.Categories.Manga()
	case "manga_hentai":
		return animelayer.Categories.MangaHentai()
	case "dorama":
		return animelayer.Categories.Dorama()
	case "music":
		return animelayer.Categories.Music()
	}
	panic("incorrect string")
}

func StrToCategoryModel(str string) model.Category {
	switch str {
	case "anime":
		return model.Categories.Anime
	case "anime_hentai":
		return model.Categories.AnimeHentai
	case "manga":
		return model.Categories.Manga
	case "manga_hentai":
		return model.Categories.MangaHentai
	case "dorama":
		return model.Categories.Dorama
	case "music":
		return model.Categories.Music
	}
	panic("incorrect string")
}
