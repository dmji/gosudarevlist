package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"
	"github.com/jackc/pgx/v5"
)

func (repo *repository) InsertUpdateNote(ctx context.Context, params model.UpdateItem) error {

	item, err := repo.GetItemByIdentifier(ctx, params.Identifier)
	if err != nil {
		return err
	}

	pgxDate, err := timeToPgTimestamp(params.Date)
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	r := repo.query.WithTx(tx)

	updateId, err := r.InsertUpdateNote(ctx, pgx_sqlc.InsertUpdateNoteParams{
		ItemID:     item.Id,
		UpdateDate: pgxDate,
		Status:     updateStatusToPgxUpdateStatus(ctx, params.UpdateStatus),
	})
	if err != nil {
		return err
	}

	updateNotes := make([]pgx_sqlc.InsertUpdateNoteItemsParams, 0, len(params.Notes))
	for _, note := range params.Notes {
		updateNotes = append(updateNotes, pgx_sqlc.InsertUpdateNoteItemsParams{
			UpdateID: updateId,
			Title:    note.ValueTitle,
			ValueOld: note.ValueOld,
			ValueNew: note.ValueNew,
		})
	}

	_, err = r.InsertUpdateNoteItems(ctx, updateNotes)

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
