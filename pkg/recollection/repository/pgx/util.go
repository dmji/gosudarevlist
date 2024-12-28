package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"
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

func statusToPgxStatus(statuses model.Status) bool {
	return statuses == model.Statuses.Completed
}

func statusesToPgxStatuses(statuses []model.Status) []bool {

	res := make([]bool, 0, len(statuses))

	for _, status := range statuses {
		res = append(res, statusToPgxStatus(status))
	}

	if len(res) == 0 {
		res = append(res, true)
		res = append(res, false)
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

func updateStatusToPgxUpdateStatus(ctx context.Context, status model.UpdateStatus) pgx_sqlc.UpdateStatus {
	switch status {
	case model.StatusNew:
		return pgx_sqlc.UpdateStatusNew
	case model.StatusRemoved:
		return pgx_sqlc.UpdateStatusRemoved
	case model.StatusUpdated:
		return pgx_sqlc.UpdateStatusUpdate
	case model.StatusUnknown:
		fallthrough
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return pgx_sqlc.UpdateStatusNew
	}
}
