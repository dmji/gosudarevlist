package repository_pgx

import (
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"time"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *repository) InsertItem(ctx context.Context, item *animelayer.ItemDetailed) error {

	lastCheckedDate := pgtype.Timestamp{}
	if err := lastCheckedDate.Scan(time.Now()); err != nil {
		return err
	}

	createdDate := pgtype.Timestamp{}
	if item.CreatedDate != nil {
		if err := createdDate.Scan(*item.CreatedDate); err != nil {
			return err
		}
	}

	updatedDate := pgtype.Timestamp{}
	if item.UpdatedDate != nil {
		if err := updatedDate.Scan(*item.UpdatedDate); err != nil {
			return err
		}
	}

	return r.query.InsertItem(ctx,
		pgx_sqlc.InsertItemParams{
			Identifier:       item.Identifier,
			Title:            item.Title,
			IsCompleted:      item.IsCompleted,
			LastCheckedDate:  lastCheckedDate,
			CreatedDate:      createdDate,
			UpdatedDate:      updatedDate,
			RefImageCover:    item.RefImageCover,
			RefImagePreview:  item.RefImagePreview,
			BlobImageCover:   "",
			BlobImagePreview: "",
			TorrentFilesSize: item.TorrentFilesSize,
			Notes:            item.Notes,
		},
	)

}
