package repository_pgx

import (
	"context"
	"log"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"
)

func (r *repository) GetUpdates(ctx context.Context, opt model.OptionsGetItems) ([]model.UpdateItem, error) {

	startID := (opt.PageIndex - 1) * opt.CountForOnePage

	items, err := r.query.GetUpdates(ctx, pgx_sqlc.GetUpdatesParams{
		Count:       int32(opt.CountForOnePage),
		OffsetCount: int32(startID),

		SearchQuery:   opt.SearchQuery,
		CategoryArray: categoriesToAnimelayerCategories(opt.Categories),
		StatusArray:   statusesToPgxStatuses(opt.Statuses),
	})

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetUpdates", "error", err)
		return nil, err
	}

	log.Printf("In-Memory repo | GetUpdates result items: %d", len(items))

	cardItems := make([]model.UpdateItem, 0, len(items))
	for _, item := range items {
		cardItems = append(cardItems, model.UpdateItem{
			Title:        item.Title,
			UpdateStatus: pgxStatusToStatus(item.UpdateStatus),
			Date:         timeFromPgTimestamp(item.UpdateDate),
		})
	}

	return cardItems, nil
}

func pgxStatusToStatus(status pgx_sqlc.UpdateStatus) model.UpdateStatus {
	switch status {
	case pgx_sqlc.UpdateStatusNew:
		return model.StatusNew
	case pgx_sqlc.UpdateStatusUpdate:
		return model.StatusUpdated
	case pgx_sqlc.UpdateStatusRemoved:
		return model.StatusRemoved
	default:
		return model.StatusUnknown
	}
}
