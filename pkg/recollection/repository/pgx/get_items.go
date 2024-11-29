package repository_pgx

import (
	"collector/pkg/logger"
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"collector/pkg/time_ru_format.go"
	"context"
	"fmt"
	"log"
)

func categoryPresentation(ctx context.Context, s model.Category, bShow bool) string {
	if bShow {
		return s.Presentation(ctx)
	}

	return ""
}

func (r *repository) GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.ItemCartData, error) {

	startID := (opt.PageIndex - 1) * opt.CountForOnePage

	items, err := r.query.GetItems(ctx, pgx_sqlc.GetItemsParams{
		Count:       int32(opt.CountForOnePage),
		OffsetCount: int32(startID),

		SearchQuery:   opt.SearchQuery,
		CategoryArray: categoriesToAnimelayerCategories(opt.Categories),
		StatusArray:   statusesToPgxStatuses(opt.Statuses),
	})

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItems", "error", err)
		return nil, err
	}

	log.Printf("In-Memory repo | GetItems result items: %d", len(items))

	cardItems := make([]model.ItemCartData, 0, len(items))
	for _, item := range items {
		cardItems = append(cardItems, model.ItemCartData{
			Title:                item.Title,
			Image:                item.RefImageCover,
			Description:          item.Notes,
			CreatedDate:          time_ru_format.Format(timeFromPgTimestamp(item.CreatedDate)),
			UpdatedDate:          time_ru_format.Format(timeFromPgTimestamp(item.UpdatedDate)),
			TorrentWeight:        item.TorrentFilesSize,
			AnimeLayerRef:        fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.Identifier),
			CategoryPresentation: categoryPresentation(ctx, pgxCategoriesToCategory(item.Category), len(opt.Categories) != 1),
			IsCompleted:          item.IsCompleted,
		})
	}

	return cardItems, nil
}
