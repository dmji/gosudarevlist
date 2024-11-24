package repository_pgx_test

import (
	"collector/cmd/env"
	"collector/pkg/logger"
	"collector/pkg/recollection/model"
	repository_pgx "collector/pkg/recollection/repository/pgx"
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetITemsByCategory(t *testing.T) {

	env.LoadEnv(10, true)

	ctx := context.Background()

	dbConfig, err := pgxpool.ParseConfig(os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		logger.Panicw(ctx, "unable to parse connString", "error", err)
	}

	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		dt, err := conn.LoadType(ctx, "CATEGORY_ANIMELAYER")
		if err != nil {
			return err
		}

		conn.TypeMap().RegisterType(dt)

		dt, err = conn.LoadType(ctx, "_CATEGORY_ANIMELAYER")
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(dt)

		return nil
	}

	connPgx, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		panic(err)
	}

	repo := repository_pgx.New(connPgx)
	repo.GetItems(ctx, model.OptionsGetItems{})
}
