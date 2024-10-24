package repository_pgx

import (
	"collector/pkg/logger"
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"log"

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

func (r *repository) GetItems(ctx context.Context, opt model.OptionsGetItems) ([]animelayer.Item, error) {
	log.Print("Pgx repo | GetItems")

	var items []pgx_sqlc.AnimelayerItem
	var err error
	items, err = r.query.GetItems(ctx, pgx_sqlc.GetItemsParams{
		Count:       int32(opt.Count),
		OffsetCount: int32(opt.Offset),
		SearchQuery: opt.SearchQuery,
		Category:    categoriesToAnimelayerCategories(opt.Category),
	})

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItems", "error", err)
		return nil, err
	}

	log.Printf("In-Memory repo | GetItems result items: %d", len(items))

	res := make([]animelayer.Item, 0, len(items))
	for _, item := range items {
		res = append(res, animelayer.Item{
			Identifier:  item.Identifier,
			Title:       item.Title,
			IsCompleted: item.IsCompleted,

			Metrics: animelayer.ItemMetrics{
				FilesSize: item.TorrentFilesSize,
			},

			RefImagePreview: item.RefImagePreview,
			RefImageCover:   item.RefImageCover,
			Updated: animelayer.ItemUpdate{
				UpdatedDate: timeFromPgTimestamp(item.UpdatedDate),
				CreatedDate: timeFromPgTimestamp(item.CreatedDate),
			},

			Notes: item.Notes,
		})
	}

	return res, nil
}
