package animelayer_repository

import (
	sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"

	"github.com/jackc/pgx/v5"
)

type txStarter interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type sqlDriver interface {
	sqlc.DBTX
	txStarter
}

type repository struct {
	query *sqlc.Queries
	db    sqlDriver
}

func NewRepository(db sqlDriver) *repository {
	return &repository{
		db:    db,
		query: sqlc.New(db),
	}
}