package main

import (
	"context"
	"database/sql"
	"errors"
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
	flag.StringVar(&identifier, "id", "67068fe95526ca6e115b3e32", "Animelayer Identifier")
	flag.StringVar(&category, "cat", "anime", "Animelayer Category")
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

	/*
	 	err = repo.RemoveItem(ctx, identifier)
	   	if err != nil {
	   		panic(err)
	   	}
	*/

	lastCheckedDate := time.Now()

	releaseStatus := model.ReleaseStatusOnAir
	if item.IsCompleted {
		releaseStatus = model.ReleaseStatusCompleted
	} else {
		year, _ := time.ParseDuration(" 1 year")
		yearAfterUpdate := item.Updated.UpdatedDate.Add(year)
		if lastCheckedDate.After(yearAfterUpdate) {
			releaseStatus = model.ReleaseStatusIncompleted
		}
	}

	_, err = repo.GetItemByIdentifier(ctx, identifier)
	newItem := &model.AnimelayerItem{
		Identifier:       item.Identifier,
		Title:            item.Title,
		ReleaseStatus:    releaseStatus,
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
	}

	if errors.Is(err, sql.ErrNoRows) {
		err = repo.InsertItem(ctx, newItem, env.StrToCategoryModel(category))
	} else {
		err = repo.UpdateItem(ctx, newItem)
	}

	if err != nil {
		panic(err)
	}
}
