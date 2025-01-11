package repository_pgx

import (
	"context"
	"fmt"
	"slices"

	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"
	"github.com/dmji/gosudarevlist/pkg/time_formater.go"
)

func categoryPresentation(ctx context.Context, s enums.Category, bShow bool) string {
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
		SimilarityThreshold: opt.SimilarityThreshold,

		SearchQuery:   opt.SearchQuery,
		CategoryArray: categoriesToAnimelayerCategories(opt.Categories, true),
		StatusArray:   releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx, opt.Statuses, true),
	})
	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItems", "error", err)
		return nil, err
	}

	cardItems := make([]model.ItemCartData, 0, len(items))
	for _, item := range items {
		if i := slices.IndexFunc(cardItems, func(e model.ItemCartData) bool { return e.Title == item.Title }); i != -1 {
			cardItems[i].AnimeLayerRefs = append(cardItems[i].AnimeLayerRefs, model.ItemCartHrefData{
				Href: fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.Identifier),
				Text: "",
			})
			continue
		}

		cardItems = append(cardItems, model.ItemCartData{
			Title:         item.Title,
			Image:         item.RefImageCover,
			Description:   item.Notes,
			CreatedDate:   time_formater.Format(ctx, pgx_utils.TimeFromPgTimestamp(item.CreatedDate)),
			UpdatedDate:   time_formater.Format(ctx, pgx_utils.TimeFromPgTimestamp(item.UpdatedDate)),
			TorrentWeight: item.TorrentFilesSize,
			AnimeLayerRefs: []model.ItemCartHrefData{
				{
					Href: fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.Identifier),
					Text: "",
				},
			},
			CategoryPresentation: categoryPresentation(ctx, pgxCategoriesToCategory(item.Category), len(opt.Categories) != 1),
			ReleaseStatus:        pgxReleaseStatusAnimelayerToReleaseStatusAnimelayer(ctx, item.ReleaseStatus),
		})
	}

	return cardItems, nil
}
