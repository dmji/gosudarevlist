package animelayer_test

import (
	"collector/internal/animelayer"
	animelayer_item_parser "collector/internal/animelayer/item_parser"
	animelayer_pages_parser "collector/internal/animelayer/pages_parser"
	"collector/pkg/model"
	"collector/pkg/parser"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"
	"testing"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
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
func getTestCreadentials() parser.Credentials {
	return parser.Credentials{
		Login:    os.Getenv("loginAnimeLayer"),
		Password: os.Getenv("passwordAnimeLayer"),
	}
}

func TestParsePagesAnime(t *testing.T) {

	ctx := context.Background()
	items := make([]model.AnimeLayerItem, 0, 2000)

	mx := sync.Mutex{}

	i := 0
	wg, ctx := errgroup.WithContext(ctx)
	wg.SetLimit(2)

	wg.Go(func() error {
		log.Print("Main loop started")

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:

				wg.Go(
					func() error {
						select {
						case <-ctx.Done():
							return ctx.Err()
						default:
							log.Printf("Reading %d page started", i+1)

							// load Html Document
							url := animelayer.FormatUrlToItemsPage("/torrents/anime", i+1)

							client, err := parser.HttpClientWithAuth(
								animelayer.BaseUrl,
								getTestCreadentials(),
							)

							if err != nil {
								return nil
							}

							doc, err := parser.LoadHtmlDocument(client, url)
							if err != nil {
								fmt.Println("Error:", err)
							}

							// Recursive parsing
							pageItems := animelayer_pages_parser.Parse(ctx, doc)
							if len(pageItems) == 0 {
								log.Printf("Reading %d page finished with error", i+1)
								return errors.New("End of context")
							}

							mx.Lock()
							defer mx.Unlock()
							items = append(items, pageItems...)
							log.Printf("Reading %d page finished", i+1)
							return nil
						}
					})

				i++
			}
		}
	})

	_ = wg.Wait()

	gotJson, _ := json.Marshal(items)
	_ = os.WriteFile("~output_pages.json", gotJson, 0644)

	expextedJson, err := os.ReadFile("~exam_pages.json")
	if err != nil {
		t.Fatal("Error", err)
	}

	expectedItems := make([]model.AnimeLayerItem, 0, 2000)
	json.Unmarshal(expextedJson, &expectedItems)
	if !slices.EqualFunc(
		items,
		expectedItems,
		func(a, b model.AnimeLayerItem) bool {
			if a.Completed != b.Completed {
				t.Logf("Expected Completed '%v', got Completed '%v'", a.Completed, b.Completed)
				return false
			}

			if a.Name != b.Name {
				t.Logf("Expected Name '%v', got Name '%v'", a.Name, b.Name)
				return false
			}

			if a.GUID != b.GUID {
				t.Fatalf("Expected GUID '%v', got GUID '%v'", a.GUID, b.GUID)
				return false
			}
			return true
		},
	) {
		t.Fail()
	}
}

func TestParseAnimeItemList(t *testing.T) {

	ctx := context.Background()

	targets := []string{
		"66d21eeabc9bb540c9713ba4",
	}

	items := make([]model.AnimeLayerItemDescription, 0)
	for _, guid := range targets {

		// load Html Document
		url := animelayer.FormatUrlToItem(guid)

		client, err := parser.HttpClientWithAuth(
			animelayer.BaseUrl,
			getTestCreadentials(),
		)

		if err != nil {
			panic(err)
		}

		doc, err := parser.LoadHtmlDocument(client, url)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// Recursive parsing
		items = append(items, *animelayer_item_parser.Parse(ctx, doc))
	}

	rankingsJson, _ := json.Marshal(items)
	_ = os.WriteFile("~output_descriptions1.json", rankingsJson, 0644)

}
