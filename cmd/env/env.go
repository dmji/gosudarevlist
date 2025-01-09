package env

import (
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"

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

func StrToCategory(s string) animelayer.Category {
	e, err := animelayer.CategoryFromString(s)
	if err != nil {
		panic("incorrect string")
	}
	return e
}

func StrToCategoryModel(str string) model.Category {
	switch str {
	case "anime":
		return model.CategoryAnime
	case "anime_hentai":
		return model.CategoryAnimeHentai
	case "manga":
		return model.CategoryManga
	case "manga_hentai":
		return model.CategoryMangaHentai
	case "dorama":
		return model.CategoryDorama
	case "music":
		return model.CategoryMusic
	}
	panic("incorrect string")
}

func AnimelayerCategoryToModelCategory(category animelayer.Category) model.Category {
	switch category {
	case animelayer.CategoryAnime:
		return model.CategoryAnime
	case animelayer.CategoryAnimeHentai:
		return model.CategoryAnimeHentai
	case animelayer.CategoryManga:
		return model.CategoryManga
	case animelayer.CategoryMangaHentai:
		return model.CategoryMangaHentai
	case animelayer.CategoryMusic:
		return model.CategoryMusic
	case animelayer.CategoryDorama:
		return model.CategoryDorama
	default:
		return model.CategoryAnime
	}
}
