package repository_pgx

import (
	"collector/pkg/logger"
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"log"
	"slices"
)

func (r *repository) GetFilters(ctx context.Context, opt model.OptionsGetItems) ([]model.FilterGroup, error) {

	isCompletedPgx, err := boolExToPgxBool(opt.IsCompleted)
	if err != nil {
		return nil, err
	}

	items, err := r.query.GetFilters(ctx, pgx_sqlc.GetFiltersParams{
		SearchQuery:   opt.SearchQuery,
		CategoryArray: categoriesToAnimelayerCategories(opt.Categories),
		IsCompleted:   isCompletedPgx,
	})

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItems", "error", err)
		return nil, err
	}

	log.Printf("In-Memory repo | GetItems result items: %d", len(items))

	cardItems := make([]model.FilterGroup, 0, 5)
	for _, item := range items {
		i := slices.IndexFunc(cardItems, func(e model.FilterGroup) bool { return e.Name == item.Name })
		if i == -1 {
			cardItems = append(cardItems, model.FilterGroup{Name: item.Name})
			i = len(cardItems) - 1
		}

		cardItems[i].Items = append(cardItems[i].Items,
			model.FilterItem{
				Value: item.Value,
				Count: item.Count,
			},
		)
	}

	return cardItems, nil
}
