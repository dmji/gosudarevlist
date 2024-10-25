package repository_pgx

import (
	"collector/internal/custom_types"
	sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func timeToPgTimestamp(t *time.Time) (pgtype.Timestamp, error) {

	pgTime := pgtype.Timestamp{}

	if t != nil {
		if err := pgTime.Scan(*t); err != nil {
			return pgTime, err
		}
	}

	return pgTime, nil
}

func timeFromPgTimestamp(t pgtype.Timestamp) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func boolExToPgxBool(b custom_types.BoolEx) (pgtype.Bool, error) {

	pgBool := pgtype.Bool{}

	switch b {
	case custom_types.BoolExTrue:
		err := pgBool.Scan(true)
		if err != nil {
			return pgBool, err
		}
	case custom_types.BoolExIntermediate:
		err := pgBool.Scan(false)
		if err != nil {
			return pgBool, err
		}
	default:
	}

	return pgBool, nil
}
