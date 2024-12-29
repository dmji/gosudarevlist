package repository_pgx

import (
	"context"
	"time"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"

	"github.com/jackc/pgx/v5"
)

func (repo *repository) UpdateItem(ctx context.Context, item *model.AnimelayerItem) error {

	now := time.Now()
	oldItem, err := repo.GetItemByIdentifier(ctx, item.Identifier)
	_ = oldItem

	//
	// Collect updated notes
	//

	notes := make([]model.UpdateItemNote, 0, 10)
	arg := pgx_sqlc.UpdateItemParams{
		/* 		Title:            pgtype.Text{},
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
		   		Identifier:       item.Identifier, */
	}

	if oldItem.Title != item.Title {
		arg.Title.Scan(item.Title)
		notes = append(notes, model.UpdateItemNote{
			ValueTitle: "Title",
			ValueOld:   oldItem.Title,
			ValueNew:   item.Title,
		})

	}

	/* 	lastCheckedDate, err := timeToPgTimestamp(&now)
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
	   	} */

	//
	// Update item
	//
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	r := repo.query.WithTx(tx)

	itemId, err := r.UpdateItem(ctx, arg)

	if err != nil {
		return err
	}

	//
	// Push updated notes
	//

	err = repo.InsertUpdateNote(ctx, model.UpdateItem{
		Date:         &now,
		UpdateStatus: model.StatusUpdated,
		Notes:        notes,
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
