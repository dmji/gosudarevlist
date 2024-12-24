package main

import (
	"context"
	"flag"
	"os"

	"github.com/dmji/gosudarevlist/cmd/env"
	"github.com/dmji/gosudarevlist/pkg/logger"
	repository_pgx "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5"
)

func init() {
	env.LoadEnv(10, true)
}

func main() {
	var category string
	flag.StringVar(&category, "cat", "", "Animelayer Category")
	flag.Parse()

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

	p := animelayer.New(animelayer.NewClientWrapper(animelayer_client))

	//
	// Repositories
	//
	connPgx, err := pgx.Connect(context.Background(), os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		panic(err)
	}

	repo := repository_pgx.New(connPgx)

	//
	// Actions
	//
	for page := 0; ; page++ {

		items, err := p.GetItemsFromCategoryPages(ctx, env.StrToCategory(category), page)
		if err != nil {
			logger.Errorw(ctx, "failed category parsing", "category", category, "page", page, "error", err)
			break
		}

		for _, item := range items {

			identifier := item.Identifier

			t, err := p.GetItemByIdentifier(ctx, identifier)
			if err != nil {
				logger.Errorw(ctx, "failed item parsing", "identifier", identifier, "error", err)
				continue
			}

			err = repo.InsertItem(ctx, t, env.StrToCategoryModel(category))
			if err != nil {
				logger.Errorw(ctx, "failed item inserting", "identifier", identifier, "error", err)
				continue
			}

			logger.Infow(ctx, "item inserted", "identifier", identifier)
		}
		break
	}
}
