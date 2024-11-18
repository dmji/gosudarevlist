package repository_pgx

import (
	"collector/pkg/logger"
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"log"
)

func (r *repository) GetUpdates(ctx context.Context, opt model.OptionsGetUpdates) ([]model.UpdateNote, error) {

	startID := (opt.PageIndex - 1) * opt.CountForOnePage

	items, err := r.query.GetUpdates(ctx, pgx_sqlc.GetUpdatesParams{
		Count:             int32(opt.CountForOnePage),
		OffsetCount:       int32(startID),
		Category:          categoriesToAnimelayerCategories(opt.Category),
		ShowAllCategories: opt.Category == model.Categories.All,
	})

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetUpdates", "error", err)
		return nil, err
	}

	log.Printf("In-Memory repo | GetUpdates result items: %d", len(items))

	cardItems := make([]model.UpdateNote, 0, len(items))
	for _, item := range items {
		cardItems = append(cardItems, model.UpdateNote{
			ItemID:      item.ItemID,
			UpdateDate:  timeFromPgTimestamp(item.UpdateDate),
			UpdateTitle: item.Title,
			ValueOld:    item.ValueOld,
			ValueNew:    item.ValueNew,
		})
	}

	return cardItems, nil
}
