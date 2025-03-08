package repository_pgx_test

import (
	"context"
	"os"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/repository"
	repository_pgx "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx"
	"github.com/dmji/gosudarevlist/pkg/env"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRepo(ctx context.Context) (repository.AnimeLayerRepositoryDriver, context.Context) {
	sugaredLogger, err := logger.New()
	if err != nil {
		panic(err)
	}

	ctx = logger.ToContext(ctx, sugaredLogger)

	err = env.LoadEnv(".env", 10)
	if err != nil {
		panic(err)
	}
	dbConfig, err := pgxpool.ParseConfig(os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		logger.Panicw(ctx, "unable to parse connString", "error", err)
	}

	dbConfig.AfterConnect = pgx_utils.AnimelayerPostgresAfterConnectFunction()

	connPgx, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		logger.Panicw(ctx, "Initialization Postgres Pool Config", "error", err)
	}

	repo := repository_pgx.New(connPgx)

	ctx = lang.New(ctx).ToContext(ctx, lang.TagEnglish)
	return repo, ctx
}
