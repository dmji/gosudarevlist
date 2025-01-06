package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"
)

func (r *repository) GetUpdates(ctx context.Context, opt model.OptionsGetItems) ([]model.UpdateItem, error) {

	startID := (opt.PageIndex - 1) * opt.CountForOnePage

	items, err := r.query.GetUpdates(ctx, pgx_sqlc.GetUpdatesParams{
		Count:               opt.CountForOnePage,
		OffsetCount:         startID,
		SimilarityThreshold: 0.05,

		SearchQuery:   opt.SearchQuery,
		CategoryArray: categoriesToAnimelayerCategories(opt.Categories, true),
		StatusArray:   releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx, opt.Statuses, true),
	})

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetUpdates", "error", err)
		return nil, err
	}

	cardItems := make([]model.UpdateItem, 0, len(items))
	for _, item := range items {
		pgxNotes, err := r.query.GetUpdateNotes(ctx, item.UpdateID)
		if err != nil {
			logger.Errorw(ctx, "Pgx repo error | GetUpdateNotes", "error", err)
			return nil, err
		}

		notes := make([]model.UpdateItemNote, 0, len(pgxNotes))
		for _, pgxNote := range pgxNotes {

			field, err := model.UpdateableFieldFromString(pgxNote.Title)
			if err != nil {
				logger.Errorw(ctx, "Pgx repo error | GetUpdateNotes string to field", "error", err)
				return nil, err
			}

			notes = append(notes, model.UpdateItemNote{
				ValueTitle: field,
				ValueOld:   pgxNote.ValueOld,
				ValueNew:   pgxNote.ValueNew,
			})
		}

		cardItems = append(cardItems, model.UpdateItem{
			Title:        item.ItemTitle,
			UpdateStatus: pgxStatusToStatus(item.UpdateStatus),
			Date:         timeFromPgTimestamp(item.UpdateDate),
			Notes:        notes,
		})
	}

	return cardItems, nil
}

func pgxStatusToStatus(status pgx_sqlc.UpdateStatus) model.UpdateStatus {
	switch status {
	case pgx_sqlc.UpdateStatusNew:
		return model.UpdateStatusNew
	case pgx_sqlc.UpdateStatusUpdate:
		return model.UpdateStatusUpdated
	case pgx_sqlc.UpdateStatusRemoved:
		return model.UpdateStatusRemoved
	default:
		return model.UpdateStatusUnknown
	}
}
