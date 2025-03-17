package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/dmji/go-myanimelist/mal"
	"github.com/dmji/gosudarevlist/assets"
	"github.com/dmji/gosudarevlist/handlers"
	"github.com/dmji/gosudarevlist/internal/animelayer_client"
	service_mal "github.com/dmji/gosudarevlist/internal/mal/service"
	repository_presenter_pgx "github.com/dmji/gosudarevlist/internal/presenter/repository/pgx"
	service_presenter "github.com/dmji/gosudarevlist/internal/presenter/service"
	service_update_manager "github.com/dmji/gosudarevlist/internal/update_manager/service"
	"github.com/dmji/gosudarevlist/internal/updater/model"
	"github.com/dmji/gosudarevlist/internal/updater/repository"
	repository_updater_pgx "github.com/dmji/gosudarevlist/internal/updater/repository/pgx"
	service_updater "github.com/dmji/gosudarevlist/internal/updater/service"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/env"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"
	"github.com/go-co-op/gocron/v2"

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
	_, bMalClientID := os.LookupEnv("MAL_CLIENT_ID")

	if !bGoose || !bLogin || !bPassword || !bMalClientID {
		err := env.LoadEnv(".env", 10)
		if err != nil {
			panic(err)
		}
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
	updateManagerService := service_update_manager.New(repoUpdater)
	updaterService := service_updater.New(repoUpdater, animelayer_client.New(animelayer_parser), updateManagerService)
	go runScheduler(ctx, updaterService)

	presentService := service_presenter.New(repoPresenter)

	//
	// Init MAL Api Client
	//
	publicInfoClient := &http.Client{
		// Create client ID from https://myanimelist.net/apiconfig.
		Transport: &clientIDTransport{ClientID: os.Getenv("MAL_CLIENT_ID")},
	}
	malApiClient, err := mal.NewSite(publicInfoClient, nil)
	if err != nil {
		logger.Panicw(ctx, "Initialization MAL Api Client", "error", err)
	}

	malService := service_mal.New(malApiClient)

	//
	// Init Router
	//
	r := handlers.New(ctx, presentService, updateManagerService, malService)
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

type clientIDTransport struct {
	Transport http.RoundTripper
	ClientID  string
}

func (c *clientIDTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.Transport == nil {
		c.Transport = http.DefaultTransport
	}
	req.Header.Add("X-MAL-CLIENT-ID", c.ClientID)
	return c.Transport.RoundTrip(req)
}

type updaterServiceClient interface {
	UpdateItemsFromCategory(ctx context.Context, category enums.Category, mode model.CategoryUpdateMode) error
}

func runScheduler(ctx context.Context, updaterService updaterServiceClient) {
	s, err := gocron.NewScheduler()
	if err != nil {
		logger.Panicw(ctx, "Scheduler creation failed", "error", err)
	}

	fnUpdate := func(ctx context.Context, cat enums.Category) {
		err = updaterService.UpdateItemsFromCategory(ctx, cat, model.CategoryUpdateModeWhileNew)
		if _, ok := repository.IsErrorItemNotChanged(err); ok {
			logger.Errorw(ctx, "RunUpdaterHandler failed", "error", err)
			return
		}
		logger.Infow(ctx, "RunUpdaterHandler completed", "category", cat)
	}

	plan := []struct {
		cron string
		cat  enums.Category
	}{
		{"0 * * * *", enums.CategoryAnime},
		{"0 */8 * * *", enums.CategoryManga},
		{"0 */8 * * *", enums.CategoryDorama},
		{"0 */8 * * *", enums.CategoryMusic},
	}

	for _, p := range plan {
		_, err = s.NewJob(
			gocron.CronJob(
				p.cron,
				false,
			),
			gocron.NewTask(
				fnUpdate,
				ctx,
				p.cat,
			),
		)
		if err != nil {
			logger.Panicw(ctx, "Scheduler Job creation failed", "error", err)
		}
	}

	s.Start()

	// TODO: add context cancel exit
	select {}

	// s.Shutdown()
}
