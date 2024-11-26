package repository_pgx

import (
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"time"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5"
)

func (repo *repository) InsertItem(ctx context.Context, item *animelayer.Item, category model.Category) error {

	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	r := repo.query.WithTx(tx)

	now := time.Now()
	lastCheckedDate, err := timeToPgTimestamp(&now)
	if err != nil {
		return err
	}

	createdDate, err := timeToPgTimestamp(item.Updated.CreatedDate)
	if err != nil {
		return err
	}

	updatedDate, err := timeToPgTimestamp(item.Updated.UpdatedDate)
	if err != nil {
		return err
	}

	itemId, err := r.InsertItem(ctx,
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
			TorrentFilesSize: item.Metrics.FilesSize,
			Notes:            item.Notes,
			Category:         categoriesToAnimelayerCategory(category),
		},
	)

	if err != nil {
		return err
	}

	err = r.InsertUpdateNote(ctx, pgx_sqlc.InsertUpdateNoteParams{
		ItemID:     itemId,
		UpdateDate: lastCheckedDate,
		//Status:     model.StatusNew,
	})

	//
	// Commit
	//
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
