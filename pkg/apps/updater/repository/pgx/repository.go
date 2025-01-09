package repository_pgx

import (
	"context"

	sqlc "github.com/dmji/gosudarevlist/pkg/apps/updater/repository/pgx/sqlc"

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

func loadType(ctx context.Context, conn *pgx.Conn, typeName string, bArray bool) error {
	dt, err := conn.LoadType(ctx, typeName)
	if err != nil {
		return err
	}

	conn.TypeMap().RegisterType(dt)

	if bArray {
		dt, err = conn.LoadType(ctx, "_"+typeName)
		if err != nil {
			return err
		}

		conn.TypeMap().RegisterType(dt)
	}

	return nil
}

func AfterConnectFunction() func(ctx context.Context, conn *pgx.Conn) error {
	return func(ctx context.Context, conn *pgx.Conn) error {
		err := loadType(ctx, conn, "CATEGORY_ANIMELAYER", true)
		if err != nil {
			return err
		}

		err = loadType(ctx, conn, "RELEASE_STATUS_ANIMELAYER", true)
		if err != nil {
			return err
		}

		return nil
	}
}
