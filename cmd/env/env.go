package env

import (
	"github.com/dmji/gosudarevlist/pkg/enums"

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

func StrToCategoryModel(str string) enums.Category {
	e, err := enums.CategoryFromString(str)
	if err != nil {
		panic("incorrect string")
	}
	return e
}

func AnimelayerCategoryToModelCategory(category animelayer.Category) enums.Category {
	switch category {
	case animelayer.CategoryAnime:
		return enums.CategoryAnime
	case animelayer.CategoryAnimeHentai:
		return enums.CategoryAnimeHentai
	case animelayer.CategoryManga:
		return enums.CategoryManga
	case animelayer.CategoryMangaHentai:
		return enums.CategoryMangaHentai
	case animelayer.CategoryMusic:
		return enums.CategoryMusic
	case animelayer.CategoryDorama:
		return enums.CategoryDorama
	default:
		return enums.CategoryAnime
	}
}
