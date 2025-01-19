package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/dmji/gosudarevlist/assets"
	"github.com/dmji/gosudarevlist/handlers"
	"github.com/dmji/gosudarevlist/internal/animelayer_client"
	"github.com/dmji/gosudarevlist/internal/env"
	repository_presenter_pgx "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx"
	service_presenter "github.com/dmji/gosudarevlist/pkg/apps/presenter/service"
	repository_updater_pgx "github.com/dmji/gosudarevlist/pkg/apps/updater/repository/pgx"
	service_updater "github.com/dmji/gosudarevlist/pkg/apps/updater/service"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"

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
		logger.Panicw(ctx, "Initialization AnimeLayer Parser", "error", err)
	}

	animelayer_parser := animelayer.New(animelayer.NewClientWrapper(animelayerClient))

	//
	// Init Service
	//
	dbConfig, err := pgxpool.ParseConfig(os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		logger.Panicw(ctx, "unable to parse connString", "error", err)
	}

	dbConfig.AfterConnect = pgx_utils.AnimelayerPostgresAfterConnectFunction()

	connPgx, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		logger.Panicw(ctx, "Initialization Postgres Pool Config", "error", err)
	}

	repoPresenter := repository_presenter_pgx.New(connPgx)
	repoUpdater := repository_updater_pgx.New(connPgx)
	updaterService := service_updater.New(repoUpdater, animelayer_client.New(animelayer_parser))
	presentService := service_presenter.New(repoPresenter)

	//
	// Init Router
	//
	r := handlers.New(ctx, presentService, updaterService)
	mux := http.NewServeMux()

	r.InitMuxWithDefaultPages(mux.HandleFunc)
	r.InitMuxWithDefaultApi(mux.HandleFunc)

	// static assets
	mux.Handle("/assets/", http.StripPrefix("/assets/", assets.Handler()))

	//
	// starting
	//
	logger.Infow(ctx, "Server starting", "port", parameter.ListenPortTcp)

	conn, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", parameter.ListenPortTcp))
	if err != nil {
		logger.Fatalw(ctx, "announces listen", "error", err)
	}

	srv := &http.Server{
		Handler:     mux,
		BaseContext: func(l net.Listener) context.Context { return ctx },
	}
	if err := srv.Serve(conn); err != nil {
		logger.Fatalw(ctx, "serve", "error", err)
	}
}
