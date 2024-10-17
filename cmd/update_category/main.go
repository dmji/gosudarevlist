package main

import (
	"collector/pkg/logger"
	repository_pgx "collector/pkg/recollection/repository/pgx"
	"context"
	"flag"
	"os"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {

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

func strToCategory(str string) animelayer.Category {
	switch str {
	case "anime":
		return animelayer.Categories.Anime()
	case "anime_hentai":
		return animelayer.Categories.AnimeHentai()
	case "manga":
		return animelayer.Categories.Manga()
	case "manga_hentai":
		return animelayer.Categories.MangaHentai()
	case "dorama":
		return animelayer.Categories.Dorama()
	case "music":
		return animelayer.Categories.Music()
	}
	panic("incorrect string")
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

		items, err := p.GetItemsFromCategoryPages(ctx, strToCategory(category), page)
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

			err = repo.InsertItem(ctx, t, strToCategory(category))
			if err != nil {
				logger.Errorw(ctx, "failed item inserting", "identifier", identifier, "error", err)
				continue
			}

			logger.Infow(ctx, "item inserted", "identifier", identifier)
		}
		break
	}
}
