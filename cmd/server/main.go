package main

import (
	"collector/handlers"
	"collector/pkg/logger"
	"collector/pkg/middleware"
	repository_pgx "collector/pkg/recollection/repository/pgx"
	"collector/pkg/recollection/service"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
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

	path := ".env"
	for i := range 10 {
		if i != 0 {
			path = "../" + path
		}
		err := godotenv.Load(path)
		if err == nil {
			return
		}
	}
	panic(".env not found")
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

	// pages
	mux.HandleFunc("/", r.HomePageHandler)
	mux.HandleFunc("/profile", r.ProfilePageHandler)
	mux.HandleFunc("/anime",
		middleware.HxPushUrlMiddleware(r.ShelfPageHandler),
	)

	// parsers
	mux.HandleFunc("/parser/animelayer", r.ScannerPageHandler)

	// api
	mux.HandleFunc("GET /api/push_url",
		middleware.HxPushUrlMiddleware(func(w http.ResponseWriter, r *http.Request) { log.Print("Changed!") }),
	)
	mux.HandleFunc("/api/cards",
		//middleware.PushQueryToRequestMiddleware(
		middleware.HxPushUrlMiddleware(
			r.ApiCards,
		),
		//),
	)
	mux.HandleFunc("/api/image", r.ApiImage)

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
