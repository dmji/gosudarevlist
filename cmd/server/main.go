package main

import (
	"collector/cmd/env"
	"collector/handlers"
	"collector/pkg/logger"
	repository_pgx "collector/pkg/recollection/repository/pgx"
	"collector/pkg/recollection/service"
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5"
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

	env.LoadEnv(10, true)
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

	//
	// Init Service
	//
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
	s := service.New(repo, animelayer_parser)
	r := handlers.New(s)

	//
	// Init Router
	//
	mux := http.NewServeMux()

	r.InitMuxWithDefaultPages(mux.HandleFunc)
	r.InitMuxWithDefaultApi(mux.HandleFunc)

	//mux.HandleFunc("/api/parser/animelayer/category", r.ApiMyAnimeListParseCategory)
	//mux.HandleFunc("/api/parser/animelayer/page", r.ApiMyAnimeListParsePage)

	// static assets
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//
	//
	// starting
	//
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
