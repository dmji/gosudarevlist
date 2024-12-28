package repository_pgx

import (
	"context"
	"time"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (repo *repository) UpdateItem(ctx context.Context, item *model.AnimelayerItem) error {

	oldItem, err := repo.GetItemByIdentifier(ctx, item.Identifier)
	_ = oldItem

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

	createdDate, err := timeToPgTimestamp(item.CreatedDate)
	if err != nil {
		return err
	}

	updatedDate, err := timeToPgTimestamp(item.UpdatedDate)
	if err != nil {
		return err
	}

	itemId, err := r.UpdateItem(ctx,
		pgx_sqlc.UpdateItemParams{
			Title:            pgtype.Text{},
			IsCompleted:      pgtype.Bool{},
			LastCheckedDate:  lastCheckedDate,
			CreatedDate:      createdDate,
			UpdatedDate:      updatedDate,
			RefImageCover:    pgtype.Text{},
			RefImagePreview:  pgtype.Text{},
			BlobImageCover:   pgtype.Text{},
			BlobImagePreview: pgtype.Text{},
			TorrentFilesSize: pgtype.Text{},
			Notes:            pgtype.Text{},
			Identifier:       item.Identifier,
		},
	)

	if err != nil {
		return err
	}

	err = r.InsertUpdateNote(ctx, pgx_sqlc.InsertUpdateNoteParams{
		ItemID:     itemId,
		UpdateDate: lastCheckedDate,
		Status:     updateStatusToPgxUpdateStatus(ctx, model.StatusUpdated),
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
