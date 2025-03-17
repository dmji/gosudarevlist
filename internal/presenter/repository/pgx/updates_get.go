package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/internal/presenter/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/internal/presenter/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"
)

func (r *repository) GetUpdates(ctx context.Context, opt model.OptionsGetItems) ([]model.UpdateItem, error) {
	startID := (opt.PageIndex - 1) * opt.CountForOnePage

	items, err := r.query.GetUpdates(ctx, pgx_sqlc.GetUpdatesParams{
		Count:               opt.CountForOnePage,
		OffsetCount:         startID,
		SimilarityThreshold: opt.SimilarityThreshold,

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

			field, err := enums.UpdateableFieldFromString(pgxNote.Title)
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
			Date:         pgx_utils.TimeFromPgTimestamp(item.UpdateDate),
			Notes:        notes,
		})
	}

	return cardItems, nil
}
