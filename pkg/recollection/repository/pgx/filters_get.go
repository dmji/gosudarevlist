package repository_pgx

import (
	"context"
	"slices"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"
)

func (r *repository) GetFilters(ctx context.Context, opt model.OptionsGetItems) ([]model.FilterGroup, error) {

	items, err := r.query.GetFilters(ctx, pgx_sqlc.GetFiltersParams{
		//SearchQuery:   opt.SearchQuery,
		CategoryArray: categoriesToAnimelayerCategories(opt.Categories),
		StatusArray:   releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx, opt.Statuses),
	})

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItems", "error", err)
		return nil, err
	}

	cardItems := make([]model.FilterGroup, 0, 5)
	for _, item := range items {
		i := slices.IndexFunc(cardItems, func(e model.FilterGroup) bool { return e.Name == item.Name })
		if i == -1 {
			cardItems = append(cardItems, model.FilterGroup{
				DisplayTitle: r.filtersStringer.GetTitlePresentation(ctx, item.Name),
				Name:         item.Name,
			})
			i = len(cardItems) - 1
		}

		cardItems[i].CheckboxItems = append(cardItems[i].CheckboxItems,
			model.FilterItem{
				Presentation: r.filtersStringer.GetItemPresentation(ctx, item.Name, item.Value),
				Value:        item.Value,
				Count:        item.Count,
			},
		)
	}

	return cardItems, nil
}
