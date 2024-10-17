package repository_pgx

import (
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"log"

	"github.com/dmji/go-animelayer-parser"
)

func (r *repository) GetItems(ctx context.Context, opt model.OptionsGetItems) ([]animelayer.Item, error) {
	log.Print("Pgx repo | GetItems")

	var items []pgx_sqlc.AnimelayerItem
	var err error
	if len(opt.SearchQuery) == 0 {
		items, err = r.query.GetItems(ctx, pgx_sqlc.GetItemsParams{
			Count:       int32(opt.Count),
			OffsetCount: int32(opt.Offset),
		})
	} else {
		items, err = r.query.GetItemsWithSearch(ctx, pgx_sqlc.GetItemsWithSearchParams{
			Count:       int32(opt.Count),
			OffsetCount: int32(opt.Offset),
			SearchQuery: opt.SearchQuery,
		})
	}

	if err != nil {
		return nil, err
	}

	log.Printf("In-Memory repo | GetItems result items: %d", len(items))

	res := make([]animelayer.Item, 0, len(items))
	for _, item := range items {
		res = append(res, animelayer.Item{
			Identifier:  item.Identifier,
			Title:       item.Title,
			IsCompleted: item.IsCompleted,

			Metrics: animelayer.ItemMetrics{
				FilesSize: item.TorrentFilesSize,
			},

			RefImagePreview: item.RefImagePreview,
			RefImageCover:   item.RefImageCover,
			Updated: animelayer.ItemUpdate{
				UpdatedDate: timeFromPgTimestamp(item.UpdatedDate),
				CreatedDate: timeFromPgTimestamp(item.CreatedDate),
			},

			Notes: item.Notes,
		})
	}

	return res, nil
}
