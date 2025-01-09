package repository_pgx_test

import (
	"context"
	"os"

	"github.com/dmji/gosudarevlist/cmd/env"
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/repository"
	repository_pgx "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx"
	"github.com/dmji/gosudarevlist/pkg/logger"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRepo(ctx context.Context) (repository.AnimeLayerRepositoryDriver, context.Context) {
	sugaredLogger, err := logger.New()
	if err != nil {
		panic(err)
	}

	ctx = logger.ToContext(ctx, sugaredLogger)

	env.LoadEnv(10, true)
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
	return repo, ctx
}
