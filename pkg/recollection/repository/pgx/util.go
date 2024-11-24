package repository_pgx

import (
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"errors"

	"github.com/dmji/go-animelayer-parser"
)

var allCategories = []pgx_sqlc.CategoryAnimelayer{
	pgx_sqlc.CategoryAnimelayerAnime,
	pgx_sqlc.CategoryAnimelayerAnimeHentai,
	pgx_sqlc.CategoryAnimelayerManga,
	pgx_sqlc.CategoryAnimelayerMangaHentai,
	pgx_sqlc.CategoryAnimelayerMusic,
	pgx_sqlc.CategoryAnimelayerDorama,
}

func categoriesToAnimelayerCategories(categories []model.Category) []pgx_sqlc.CategoryAnimelayer {
	res := make([]pgx_sqlc.CategoryAnimelayer, 0, len(categories))

	for _, category := range categories {
		res = append(res, categoriesToAnimelayerCategory(category))
	}

	if len(res) == 0 {
		for _, cat := range allCategories {
			res = append(res, cat)
		}
	}

	return res
}

func categoriesToAnimelayerCategory(category model.Category) pgx_sqlc.CategoryAnimelayer {

	switch category {
	case model.Categories.Anime:
		return pgx_sqlc.CategoryAnimelayerAnime
	case model.Categories.AnimeHentai:
		return pgx_sqlc.CategoryAnimelayerAnimeHentai
	case model.Categories.Manga:
		return pgx_sqlc.CategoryAnimelayerManga
	case model.Categories.MangaHentai:
		return pgx_sqlc.CategoryAnimelayerMangaHentai
	case model.Categories.Music:
		return pgx_sqlc.CategoryAnimelayerMusic
	case model.Categories.Dorama:
		return pgx_sqlc.CategoryAnimelayerDorama
	default:
		return pgx_sqlc.CategoryAnimelayerAnime
	}
}

func pgxCategoriesToCategory(category pgx_sqlc.CategoryAnimelayer) model.Category {

	switch category {
	case pgx_sqlc.CategoryAnimelayerAnime:
		return model.Categories.Anime
	case pgx_sqlc.CategoryAnimelayerAnimeHentai:
		return model.Categories.AnimeHentai
	case pgx_sqlc.CategoryAnimelayerManga:
		return model.Categories.Manga
	case pgx_sqlc.CategoryAnimelayerMangaHentai:
		return model.Categories.MangaHentai
	case pgx_sqlc.CategoryAnimelayerMusic:
		return model.Categories.Music
	case pgx_sqlc.CategoryAnimelayerDorama:
		return model.Categories.Dorama
	default:
		return model.Categories.Anime
	}
}

func categoryToPgxCategory(cat animelayer.Category) (pgx_sqlc.CategoryAnimelayer, error) {
	switch cat {
	case animelayer.Categories.Anime():
		return pgx_sqlc.CategoryAnimelayerAnime, nil
	case animelayer.Categories.AnimeHentai():
		return pgx_sqlc.CategoryAnimelayerAnimeHentai, nil
	case animelayer.Categories.Manga():
		return pgx_sqlc.CategoryAnimelayerManga, nil
	case animelayer.Categories.MangaHentai():
		return pgx_sqlc.CategoryAnimelayerMangaHentai, nil
	case animelayer.Categories.Dorama():
		return pgx_sqlc.CategoryAnimelayerDorama, nil
	case animelayer.Categories.Music():
		return pgx_sqlc.CategoryAnimelayerMusic, nil
	}

	return pgx_sqlc.CategoryAnimelayerAnime, errors.New("undefined category")
}
