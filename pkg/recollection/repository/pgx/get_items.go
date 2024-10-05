package repository_pgx

import (
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"log"
	"time"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5/pgtype"
)

func timeFromPgTimestamp(t pgtype.Timestamp) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func (r *repository) GetItems(ctx context.Context, opt model.OptionsGetItems) ([]animelayer.ItemDetailed, error) {
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

	res := make([]animelayer.ItemDetailed, 0, len(items))
	for _, item := range items {
		res = append(res, animelayer.ItemDetailed{
			Identifier:  item.Identifier,
			Title:       item.Title,
			IsCompleted: item.IsCompleted,

			TorrentFilesSize: item.TorrentFilesSize,

			RefImagePreview: item.RefImagePreview,
			RefImageCover:   item.RefImageCover,

			UpdatedDate: timeFromPgTimestamp(item.UpdatedDate),
			CreatedDate: timeFromPgTimestamp(item.CreatedDate),
		})
	}

	return res, nil
}
