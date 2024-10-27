package repository_pgx

import (
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"errors"

	"github.com/dmji/go-animelayer-parser"
)

/* func categoriesToAnimelayerCategories(categories []model.Category) pgx_sqlc.CategoryAnimelayer {
	res := make([]pgx_sqlc.CategoryAnimelayer, 0, len(categories))

	for _, category := range categories {
		switch category {
		case model.Categories.Anime:
			res = append(res, pgx_sqlc.CategoryAnimelayerAnime)
		case model.Categories.AnimeHentai:
			res = append(res, pgx_sqlc.CategoryAnimelayerAnimeHentai)
		case model.Categories.Manga:
			res = append(res, pgx_sqlc.CategoryAnimelayerManga)
		case model.Categories.MangaHentai:
			res = append(res, pgx_sqlc.CategoryAnimelayerMangaHentai)
		case model.Categories.Music:
			res = append(res, pgx_sqlc.CategoryAnimelayerMusic)
		case model.Categories.Dorama:
			res = append(res, pgx_sqlc.CategoryAnimelayerDorama)
		}
	}

	if len(res) == 0 {
		res = append(res, pgx_sqlc.CategoryAnimelayerAnime)
		res = append(res, pgx_sqlc.CategoryAnimelayerAnimeHentai)
		res = append(res, pgx_sqlc.CategoryAnimelayerManga)
		res = append(res, pgx_sqlc.CategoryAnimelayerMangaHentai)
		res = append(res, pgx_sqlc.CategoryAnimelayerMusic)
		res = append(res, pgx_sqlc.CategoryAnimelayerDorama)
	}

	return res[0]
} */

func categoriesToAnimelayerCategories(category model.Category) pgx_sqlc.CategoryAnimelayer {

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
