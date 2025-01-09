package repository_pgx

import (
	"context"

	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/apps/updater/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/enums"
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

func categoriesToAnimelayerCategories(categories []enums.Category, AppendAllOnEmpty bool) []pgx_sqlc.CategoryAnimelayer {
	res := make([]pgx_sqlc.CategoryAnimelayer, 0, len(categories))

	for _, category := range categories {
		res = append(res, categoriesToAnimelayerCategory(category))
	}

	if AppendAllOnEmpty && len(res) == 0 {
		res = append(res, allCategories...)
	}

	return res
}

func categoriesToAnimelayerCategory(category enums.Category) pgx_sqlc.CategoryAnimelayer {
	switch category {

	case enums.CategoryAnime:
		return pgx_sqlc.CategoryAnimelayerAnime
	case enums.CategoryAnimeHentai:
		return pgx_sqlc.CategoryAnimelayerAnimeHentai
	case enums.CategoryManga:
		return pgx_sqlc.CategoryAnimelayerManga
	case enums.CategoryMangaHentai:
		return pgx_sqlc.CategoryAnimelayerMangaHentai
	case enums.CategoryMusic:
		return pgx_sqlc.CategoryAnimelayerMusic
	case enums.CategoryDorama:
		return pgx_sqlc.CategoryAnimelayerDorama
	default:
		return pgx_sqlc.CategoryAnimelayerAnime
	}
}

func pgxCategoriesToCategory(category pgx_sqlc.CategoryAnimelayer) enums.Category {
	switch category {

	case pgx_sqlc.CategoryAnimelayerAnime:
		return enums.CategoryAnime
	case pgx_sqlc.CategoryAnimelayerAnimeHentai:
		return enums.CategoryAnimeHentai
	case pgx_sqlc.CategoryAnimelayerManga:
		return enums.CategoryManga
	case pgx_sqlc.CategoryAnimelayerMangaHentai:
		return enums.CategoryMangaHentai
	case pgx_sqlc.CategoryAnimelayerMusic:
		return enums.CategoryMusic
	case pgx_sqlc.CategoryAnimelayerDorama:
		return enums.CategoryDorama
	default:
		return enums.CategoryAnime
	}
}

func updateStatusToPgxUpdateStatus(ctx context.Context, status enums.UpdateStatus) pgx_sqlc.UpdateStatus {
	switch status {

	case enums.UpdateStatusNew:
		return pgx_sqlc.UpdateStatusNew
	case enums.UpdateStatusRemoved:
		return pgx_sqlc.UpdateStatusRemoved
	case enums.UpdateStatusUpdated:
		return pgx_sqlc.UpdateStatusUpdate
	case enums.UpdateStatusUnknown:
		fallthrough
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return pgx_sqlc.UpdateStatusNew
	}
}

func releaseStatusAnimelayerToPgxReleaseStatusAnimelayer(ctx context.Context, status enums.ReleaseStatus) pgx_sqlc.ReleaseStatusAnimelayer {
	switch status {

	case enums.ReleaseStatusOnAir:
		return pgx_sqlc.ReleaseStatusAnimelayerOnAir
	case enums.ReleaseStatusIncompleted:
		return pgx_sqlc.ReleaseStatusAnimelayerIncompleted
	case enums.ReleaseStatusCompleted:
		return pgx_sqlc.ReleaseStatusAnimelayerCompleted
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return pgx_sqlc.ReleaseStatusAnimelayerIncompleted
	}
}

func pgxReleaseStatusAnimelayerToReleaseStatusAnimelayer(ctx context.Context, status pgx_sqlc.ReleaseStatusAnimelayer) enums.ReleaseStatus {
	switch status {

	case pgx_sqlc.ReleaseStatusAnimelayerOnAir:
		return enums.ReleaseStatusOnAir
	case pgx_sqlc.ReleaseStatusAnimelayerIncompleted:
		return enums.ReleaseStatusIncompleted
	case pgx_sqlc.ReleaseStatusAnimelayerCompleted:
		return enums.ReleaseStatusCompleted
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return enums.ReleaseStatusIncompleted
	}
}

var allReleaseStatus = []pgx_sqlc.ReleaseStatusAnimelayer{
	pgx_sqlc.ReleaseStatusAnimelayerOnAir,
	pgx_sqlc.ReleaseStatusAnimelayerIncompleted,
	pgx_sqlc.ReleaseStatusAnimelayerCompleted,
}

func releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx context.Context, statuses []enums.ReleaseStatus, AppendAllOnEmpty bool) []pgx_sqlc.ReleaseStatusAnimelayer {
	res := make([]pgx_sqlc.ReleaseStatusAnimelayer, 0, len(statuses))

	for _, status := range statuses {
		res = append(res, releaseStatusAnimelayerToPgxReleaseStatusAnimelayer(ctx, status))
	}

	if AppendAllOnEmpty && len(res) == 0 {
		res = append(res, allReleaseStatus...)
	}

	return res
}
