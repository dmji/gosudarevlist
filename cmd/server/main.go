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
	"github.com/dmji/gosudarevlist/internal/env"
	repository_pgx "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx"
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/service"
	"github.com/dmji/gosudarevlist/pkg/logger"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5/pgxpool"
)

type cliParameters struct {
	ListenPortTcp int64
}

var parameter = cliParameters{
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
	animelayer_client, err := animelayer.DefaultClientWithAuth(animelayer_credentials)
	if err != nil {
		panic(err)
	}

	animelayer_parser := animelayer.New(animelayer.NewClientWrapper(animelayer_client))
	_ = animelayer_parser
	//
	// Init Service
	//
	dbConfig, err := pgxpool.ParseConfig(os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		logger.Panicw(ctx, "unable to parse connString", "error", err)
	}

	dbConfig.AfterConnect = repository_pgx.AfterConnectFunction()

	connPgx, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		panic(err)
	}

	repo := repository_pgx.New(connPgx)
	s := service.New(repo)
	r := handlers.New(ctx, s)

	//
	// Init Router
	//
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
