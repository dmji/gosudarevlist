package repository_pgx

import (
	"context"
	"time"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"

	"github.com/jackc/pgx/v5"
)

func (repo *repository) InsertItem(ctx context.Context, item *model.AnimelayerItem, category model.Category) error {

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
			TorrentFilesSize: item.TorrentFilesSize,
			Notes:            item.Notes,
			Category:         categoriesToAnimelayerCategory(category),
		},
	)

	if err != nil {
		return err
	}

	err = repo.InsertUpdateNote(ctx, model.UpdateItem{
		Date:         &now,
		UpdateStatus: model.StatusNew,
		Notes:        []model.UpdateItemNote{},
		ItemId:       itemId,
		//Identifier:   item.Identifier,
	})

	if err != nil {
		return err
	}

	//
	// Commit
	//
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
