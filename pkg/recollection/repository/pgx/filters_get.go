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
		CheckedCategoryArray:  categoriesToAnimelayerCategories(opt.Categories, false),
		CheckedStatusArray:    releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx, opt.Statuses, false),
		SelectedCategoryArray: categoriesToAnimelayerCategories(opt.Categories, true),
		SelectedStatusArray:   releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx, opt.Statuses, true),

		SearchQuery:         opt.SearchQuery,
		SimilarityThreshold: opt.SimilarityThreshold,
	})

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItems", "error", err)
		return nil, err
	}

	cardItems := make([]model.FilterGroup, 0, 5)
	for _, item := range items {

		filterType, err := model.FilterFromString(item.Name)
		if err != nil {
			logger.Errorw(ctx, "failed parse filter type", "error", err)
			continue
		}

		i := slices.IndexFunc(cardItems, func(e model.FilterGroup) bool { return e.Name == item.Name })
		if i == -1 {
			cardItems = append(cardItems, model.FilterGroup{
				DisplayTitle: filterType.Presentation(ctx),
				Name:         item.Name,
			})
			i = len(cardItems) - 1
		}

		present, err := filterType.ChildsPresentation(ctx, item.Value)
		if err != nil {
			logger.Errorw(ctx, "failed parse filter type child presentation", "error", err)
			continue
		}

		cardItems[i].CheckboxItems = append(cardItems[i].CheckboxItems,
			model.FilterItem{
				Presentation:  present,
				Value:         item.Value,
				Count:         item.Count,
				CountFiltered: item.CountFiltered,
				Selected:      item.Selected,
			},
		)
	}

	return cardItems, nil
}
