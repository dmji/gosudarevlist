package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/dmji/go-animelayer-parser"
	jikan "github.com/dmji/go-jikan"
	"github.com/dmji/go-jikan/models/operations"
	"github.com/dmji/go-jikan/types"
	repository_updater_pgx "github.com/dmji/gosudarevlist/internal/updater/repository/pgx"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/env"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"

	"github.com/hbakhtiyor/strsim"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/ratelimit"
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

	if !bGoose {
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
	// Init Service
	//
	dbConfig, err := pgxpool.ParseConfig(os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		logger.Panicw(ctx, "unable to parse connString", "error", err)
	}

	dbConfig.AfterConnect = pgx_utils.AnimelayerPostgresAfterConnectFunction()

	connPgx, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		panic(err)
	}

	//
	// Init Updater
	//
	repoUpdater := repository_updater_pgx.New(connPgx)
	itemsAll, err := repoUpdater.GetItems(ctx, time.Time{})
	if err != nil {
		panic(err)
	}

	type Item struct {
		identifier []string
		title      string
		year       string
		epCount    string
	}
	items := make([]Item, 0, 1000)

	for _, item := range itemsAll {
		if item.ReleaseStatus != enums.ReleaseStatusOnAir {
			continue
		}

		m := animelayer.TryGetSomthingSemantizedFromNotes(item.Notes)

		it := Item{
			identifier: []string{item.Identifier},
			title:      item.Title,
			year:       traverseMapNotesSemantized("Год выхода", m),
			epCount:    traverseMapNotesSemantized("Кол серий", m),
		}

		i := slices.IndexFunc(items,
			func(e Item) bool {
				return e.title == it.title &&
					e.year == it.year &&
					e.epCount == it.epCount
			},
		)
		if i == -1 {
			items = append(items, it)
			continue
		}

		items[i].identifier = append(items[i].identifier, item.Identifier)
	}

	//
	// Jikan Loop
	//
	jk := jikan.New()

	s := "One Piece Egghead Arc (1086—1106)"
	res, err := jk.Anime.Search(ctx, operations.GetAnimeSearchRequest{
		Q:         &s,
		StartDate: types.String(fmt.Sprintf("%s-01-01", "2025")),
		// EndDate:   types.String(fmt.Sprintf("%s-12-31", time.)),
	})
	if err != nil {
		log.Print(err)
	}
	_ = res

	_ = ratelimit.New(1, ratelimit.Per(5*time.Second)) // per second
	for _, item := range items {
		// for title := range strings.SplitSeq(item.Title, "/")
		{

			// rl.Take()
			time.Sleep(1 * time.Second)
			res, err := jk.Anime.Search(ctx, operations.GetAnimeSearchRequest{
				Q:         &item.title,
				StartDate: types.String(fmt.Sprintf("%s-01-01", item.year)),
				// EndDate:   types.String(fmt.Sprintf("%s-12-31", time.)),
			})
			if err != nil {
				log.Print(err)
				continue
			}
			as := res.GetAnimeSearch()
			if as == nil {
				log.Println("ERROR: empty search result for title: ", item.title)
				continue
			}
			_ = res

			data := as.GetData()
			if len(data) == 0 {
				log.Println("ERROR: empty data for title: ", item.title)
				continue
			}

			titles := make([]string, 0, len(data))
			for _, d := range data {
				titles = append(titles, *d.GetTitle())
			}

			matches, err := strsim.FindBestMatch(item.title, titles)
			if err != nil {
				fmt.Printf("ERROR finding best match: '%s'\n", err.Error())
				continue
			}

			if matches.BestMatch.Score < 0.25 {
				fmt.Printf("ERROR Title: '%s' no matches\n", item.title)
				continue
			}

			fmt.Printf("Title: '%s', Best match: '%s' with score=%f\n", item.title, matches.BestMatch.Target, matches.BestMatch.Score)
		}
	}
}

func traverseMapNotesSemantized(tag string, m *animelayer.NotesSematizied) string {
	for _, t := range m.Taged {
		if t.Tag == tag && len(t.Text) != 0 {
			return t.Text
		}
		if t.Childs != nil {
			s := traverseMapNotesSemantized(tag, t.Childs)
			if s != "" {
				return s
			}
		}
	}
	return ""
}

func splitGenres(genreString string) []string {
	genres := strings.Split(genreString, ",")
	for gi := range genres {
		for {
			if len(genres[gi]) == 0 {
				break
			}
			if genres[gi][0] == ' ' || genres[gi][0] == ':' {
				genres[gi] = genres[gi][1:]
				continue
			}

			l := len(genres[gi]) - 1
			if genres[gi][l] == ' ' || genres[gi][l] == ':' {
				genres[gi] = genres[gi][:l-1]
				continue
			}
			break
		}
	}
	return slices.DeleteFunc(genres, func(e string) bool { return len(e) == 0 })
}
