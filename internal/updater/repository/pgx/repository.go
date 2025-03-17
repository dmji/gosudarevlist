package repository_pgx

import (
	"context"

	sqlc "github.com/dmji/gosudarevlist/internal/updater/repository/pgx/sqlc"

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

func New(db sqlDriver) *repository {
	return &repository{
		db:    db,
		query: sqlc.New(db),
	}
}
