package main

import (
	"collector/cmd/env"
	"collector/pkg/logger"
	repository_pgx "collector/pkg/recollection/repository/pgx"
	"context"
	"errors"
	"flag"
	"os"

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
	category = "anime_hentai"

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
	for page := 1; ; page++ {

		items, err := p.GetItemsFromCategoryPages(ctx, env.StrToCategory(category), page)
		if err != nil {
			if errors.Is(err, animelayer.ErrorEmptyPage) {
				break
			}

			logger.Errorw(ctx, "failed category parsing", "category", category, "page", page, "error", err)
			break
		}

		for _, item := range items {
			err = repo.InsertItem(ctx, &item, env.StrToCategoryModel(category))
			if err != nil {
				logger.Errorw(ctx, "failed item inserting", "identifier", item.Identifier, "error", err)
				continue
			}
			logger.Infow(ctx, "item inserted", "identifier", item.Identifier)

			/*
				identifier := item.Identifier

				t, err := p.GetItemByIdentifier(ctx, identifier)
				if err != nil {
					logger.Errorw(ctx, "failed item parsing", "identifier", identifier, "error", err)
					continue
				}

				err = repo.InsertItem(ctx, t, env.StrToCategory(category))
				if err != nil {
					logger.Errorw(ctx, "failed item inserting", "identifier", identifier, "error", err)
					continue
				}
				logger.Infow(ctx, "item inserted", "identifier", identifier)
			*/
		}
	}
}
