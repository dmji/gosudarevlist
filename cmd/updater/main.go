package main

import (
	"context"
	"flag"
	"os"

	"github.com/dmji/gosudarevlist/internal/animelayer_client"
	"github.com/dmji/gosudarevlist/internal/env"
	repository_presenter_pgx "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	repository_updater_pgx "github.com/dmji/gosudarevlist/pkg/apps/updater/repository/pgx"
	service_updater "github.com/dmji/gosudarevlist/pkg/apps/updater/service"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5/pgxpool"
)

var parameter = struct {
	ListenPortTcp int64
}{
	ListenPortTcp: 8080,
}

func init() {
	flag.Int64Var(&parameter.ListenPortTcp, "port", 8080, "Port for tcp connection")
	flag.Parse()

	_, bGoose := os.LookupEnv("GOOSE_DBSTRING")
	_, bLogin := os.LookupEnv("ANIMELAYER_LOGIN")
	_, bPassword := os.LookupEnv("ANIMELAYER_PASSWORD")

	if !bGoose || !bLogin || !bPassword {
		env.LoadEnv(10, true)
	}
}

func main() {
	//
	// Init logger
	//
	sugaredLogger, err := logger.New()
	if err != nil {
		panic(err)
	}

	ctx := logger.ToContext(context.Background(), sugaredLogger)

	//
	// Init Animelayer Parser
	//
	animelayer_credentials := animelayer.Credentials{
		Login:    os.Getenv("ANIMELAYER_LOGIN"),
		Password: os.Getenv("ANIMELAYER_PASSWORD"),
	}
	animelayerClient, err := animelayer.DefaultClientWithAuth(animelayer_credentials)
	if err != nil {
		panic(err)
	}

	animelayer_parser := animelayer.New(animelayer.NewClientWrapper(animelayerClient))

	//
	// Init Service
	//
	dbConfig, err := pgxpool.ParseConfig(os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		logger.Panicw(ctx, "unable to parse connString", "error", err)
	}

	dbConfig.AfterConnect = repository_presenter_pgx.AfterConnectFunction()

	connPgx, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		panic(err)
	}

	repoUpdater := repository_updater_pgx.New(connPgx)
	updaterService := service_updater.New(repoUpdater, animelayer_client.New(animelayer_parser))

	// cat := enums.CategoryAnime
	// cat := enums.CategoryManga
	// cat := enums.CategoryDorama
	// cat := enums.CategoryMusic
	cat := enums.CategoryAll
	updaterService.UpdateItemsFromCategory(ctx, cat, model.CategoryUpdateModeAll)
}