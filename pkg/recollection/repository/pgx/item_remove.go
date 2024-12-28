package repository_pgx

import (
	"context"
	"time"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"
	"github.com/jackc/pgx/v5"
)

func (repo *repository) RemoveItem(ctx context.Context, identifier string) error {

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

	itemId, err := r.RemoveItem(ctx, identifier)
	if err != nil {
		return err
	}

	err = r.InsertUpdateNote(ctx, pgx_sqlc.InsertUpdateNoteParams{
		ItemID:     itemId,
		UpdateDate: lastCheckedDate,
		Status:     updateStatusToPgxUpdateStatus(ctx, model.StatusRemoved),
	})

	return nil
}
