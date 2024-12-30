package repository_pgx

import (
	"context"
	"fmt"
	"log"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/time_ru_format.go"
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
		Count:               opt.CountForOnePage,
		OffsetCount:         startID,
		SimilarityThreshold: 0.05,

		SearchQuery:   opt.SearchQuery,
		CategoryArray: categoriesToAnimelayerCategories(opt.Categories),
		StatusArray:   releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx, opt.Statuses),
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
			ReleaseStatus:        pgxReleaseStatusAnimelayerToReleaseStatusAnimelayer(ctx, item.ReleaseStatus),
		})
	}

	return cardItems, nil
}
