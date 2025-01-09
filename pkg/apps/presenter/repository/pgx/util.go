package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

var allCategories = []pgx_sqlc.CategoryAnimelayer{
	pgx_sqlc.CategoryAnimelayerAnime,
	pgx_sqlc.CategoryAnimelayerAnimeHentai,
	pgx_sqlc.CategoryAnimelayerManga,
	pgx_sqlc.CategoryAnimelayerMangaHentai,
	pgx_sqlc.CategoryAnimelayerMusic,
	pgx_sqlc.CategoryAnimelayerDorama,
}

func categoriesToAnimelayerCategories(categories []model.Category, AppendAllOnEmpty bool) []pgx_sqlc.CategoryAnimelayer {
	res := make([]pgx_sqlc.CategoryAnimelayer, 0, len(categories))

	for _, category := range categories {
		res = append(res, categoriesToAnimelayerCategory(category))
	}

	if AppendAllOnEmpty && len(res) == 0 {
		res = append(res, allCategories...)
	}

	return res
}

func categoriesToAnimelayerCategory(category model.Category) pgx_sqlc.CategoryAnimelayer {
	switch category {

	case model.CategoryAnime:
		return pgx_sqlc.CategoryAnimelayerAnime
	case model.CategoryAnimeHentai:
		return pgx_sqlc.CategoryAnimelayerAnimeHentai
	case model.CategoryManga:
		return pgx_sqlc.CategoryAnimelayerManga
	case model.CategoryMangaHentai:
		return pgx_sqlc.CategoryAnimelayerMangaHentai
	case model.CategoryMusic:
		return pgx_sqlc.CategoryAnimelayerMusic
	case model.CategoryDorama:
		return pgx_sqlc.CategoryAnimelayerDorama
	default:
		return pgx_sqlc.CategoryAnimelayerAnime
	}
}

func pgxCategoriesToCategory(category pgx_sqlc.CategoryAnimelayer) model.Category {
	switch category {

	case pgx_sqlc.CategoryAnimelayerAnime:
		return model.CategoryAnime
	case pgx_sqlc.CategoryAnimelayerAnimeHentai:
		return model.CategoryAnimeHentai
	case pgx_sqlc.CategoryAnimelayerManga:
		return model.CategoryManga
	case pgx_sqlc.CategoryAnimelayerMangaHentai:
		return model.CategoryMangaHentai
	case pgx_sqlc.CategoryAnimelayerMusic:
		return model.CategoryMusic
	case pgx_sqlc.CategoryAnimelayerDorama:
		return model.CategoryDorama
	default:
		return model.CategoryAnime
	}
}

func updateStatusToPgxUpdateStatus(ctx context.Context, status model.UpdateStatus) pgx_sqlc.UpdateStatus {
	switch status {

	case model.UpdateStatusNew:
		return pgx_sqlc.UpdateStatusNew
	case model.UpdateStatusRemoved:
		return pgx_sqlc.UpdateStatusRemoved
	case model.UpdateStatusUpdated:
		return pgx_sqlc.UpdateStatusUpdate
	case model.UpdateStatusUnknown:
		fallthrough
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return pgx_sqlc.UpdateStatusNew
	}
}

func releaseStatusAnimelayerToPgxReleaseStatusAnimelayer(ctx context.Context, status model.ReleaseStatus) pgx_sqlc.ReleaseStatusAnimelayer {
	switch status {

	case model.ReleaseStatusOnAir:
		return pgx_sqlc.ReleaseStatusAnimelayerOnAir
	case model.ReleaseStatusIncompleted:
		return pgx_sqlc.ReleaseStatusAnimelayerIncompleted
	case model.ReleaseStatusCompleted:
		return pgx_sqlc.ReleaseStatusAnimelayerCompleted
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return pgx_sqlc.ReleaseStatusAnimelayerIncompleted
	}
}

func pgxReleaseStatusAnimelayerToReleaseStatusAnimelayer(ctx context.Context, status pgx_sqlc.ReleaseStatusAnimelayer) model.ReleaseStatus {
	switch status {

	case pgx_sqlc.ReleaseStatusAnimelayerOnAir:
		return model.ReleaseStatusOnAir
	case pgx_sqlc.ReleaseStatusAnimelayerIncompleted:
		return model.ReleaseStatusIncompleted
	case pgx_sqlc.ReleaseStatusAnimelayerCompleted:
		return model.ReleaseStatusCompleted
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return model.ReleaseStatusIncompleted
	}
}

var allReleaseStatus = []pgx_sqlc.ReleaseStatusAnimelayer{
	pgx_sqlc.ReleaseStatusAnimelayerOnAir,
	pgx_sqlc.ReleaseStatusAnimelayerIncompleted,
	pgx_sqlc.ReleaseStatusAnimelayerCompleted,
}

func releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx context.Context, statuses []model.ReleaseStatus, AppendAllOnEmpty bool) []pgx_sqlc.ReleaseStatusAnimelayer {
	res := make([]pgx_sqlc.ReleaseStatusAnimelayer, 0, len(statuses))

	for _, status := range statuses {
		res = append(res, releaseStatusAnimelayerToPgxReleaseStatusAnimelayer(ctx, status))
	}

	if AppendAllOnEmpty && len(res) == 0 {
		res = append(res, allReleaseStatus...)
	}

	return res
}