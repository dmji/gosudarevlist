package main

import (
	"collector/cmd/env"
	"collector/pkg/logger"
	repository_pgx "collector/pkg/recollection/repository/pgx"
	"context"
	"flag"
	"os"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5"
)

func init() {
	env.LoadEnv(10, true)
}

func main() {
	var identifier, category string
	flag.StringVar(&identifier, "id", "", "Animelayer Identifier")
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
	t, err := p.GetItemByIdentifier(ctx, identifier)
	if err != nil {
		panic(err)
	}

	err = repo.RemoveItem(ctx, identifier)
	if err != nil {
		panic(err)
	}

	err = repo.InsertItem(ctx, t, env.StrToCategoryModel(category))
	if err != nil {
		panic(err)
	}
}
