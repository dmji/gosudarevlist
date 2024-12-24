package repository_pgx

import (
	"context"
	"time"

	"github.com/dmji/gosudarevlist/internal/query_cards"
	sqlc "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx/sqlc"

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
	query           *sqlc.Queries
	db              sqlDriver
	filtersStringer *query_cards.Stringer
}

func New(db sqlDriver) *repository {
	return &repository{
		db:              db,
		query:           sqlc.New(db),
		filtersStringer: query_cards.NewStringer(),
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
