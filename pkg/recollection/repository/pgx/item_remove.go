package repository_pgx

import (
	"context"
	"time"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	"github.com/jackc/pgx/v5"
)

func (repo *repository) RemoveItem(ctx context.Context, identifier string) error {

	now := time.Now()

	//
	// Start transaction
	//
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	r := repo.query.WithTx(tx)

	itemId, err := r.RemoveItem(ctx, identifier)
	if err != nil {
		return err
	}

	err = repo.InsertUpdateNote(ctx, model.UpdateItem{
		Date:         &now,
		UpdateStatus: model.StatusRemoved,
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
