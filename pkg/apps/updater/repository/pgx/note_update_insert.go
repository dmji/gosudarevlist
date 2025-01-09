package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/apps/updater/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"
	"github.com/jackc/pgx/v5"
)

func (repo *repository) InsertUpdateNote(ctx context.Context, params model.UpdateItem) error {
	itemId := params.ItemId
	if itemId == 0 {
		item, err := repo.GetItemByIdentifier(ctx, params.Identifier)
		if err != nil {
			return err
		}
		itemId = item.Id
	}

	pgxDate, err := pgx_utils.TimeToPgTimestamp(params.Date)
	if err != nil {
		return err
	}

	//
	// Start transaction
	//
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	r := repo.query.WithTx(tx)

	updateId, err := r.InsertUpdateNote(ctx, pgx_sqlc.InsertUpdateNoteParams{
		ItemID:     itemId,
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
			Title:    note.ValueTitle.String(),
			ValueOld: note.ValueOld,
			ValueNew: note.ValueNew,
		})
	}

	t, err := r.InsertUpdateNoteItems(ctx, updateNotes)
	_ = t

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
