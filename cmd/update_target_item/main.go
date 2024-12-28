package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/dmji/gosudarevlist/cmd/env"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	repository_pgx "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx"

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
	item, err := p.GetItemByIdentifier(ctx, identifier)
	if err != nil {
		panic(err)
	}

	err = repo.RemoveItem(ctx, identifier)
	if err != nil {
		panic(err)
	}

	lastCheckedDate := time.Now()

	err = repo.InsertItem(ctx, &model.AnimelayerItem{
		Identifier:       item.Identifier,
		Title:            item.Title,
		IsCompleted:      item.IsCompleted,
		LastCheckedDate:  &lastCheckedDate,
		CreatedDate:      item.Updated.CreatedDate,
		UpdatedDate:      item.Updated.UpdatedDate,
		RefImageCover:    item.RefImageCover,
		RefImagePreview:  item.RefImagePreview,
		BlobImageCover:   "",
		BlobImagePreview: "",
		TorrentFilesSize: item.Metrics.FilesSize,
		Notes:            item.Notes,
		Category:         env.AnimelayerCategoryToModelCategory(item.Category),
	}, env.StrToCategoryModel(category))
	if err != nil {
		panic(err)
	}
}
