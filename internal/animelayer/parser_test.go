package animelayer_test

import (
	animelayer_item_parser "collector/internal/animelayer/item_parser"
	animelayer_pages_parser "collector/internal/animelayer/pages_parser"
	"collector/pkg/model"
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"slices"
	"sync"
	"testing"

	"golang.org/x/sync/errgroup"
)

const (
	BASE_ADDRESS_URI = "https://animelayer.ru"
)

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

							pageItems := animelayer_pages_parser.CollectBaseItemsFromAddress(ctx, BASE_ADDRESS_URI, "/torrents/anime", i+1)
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

func TestParseAnimeItem(t *testing.T) {

	ctx := context.Background()

	item := animelayer_item_parser.CollectItemFromAddress(ctx, BASE_ADDRESS_URI, "6686e8e82f708c215170c809")

	rankingsJson, _ := json.Marshal(item)
	_ = os.WriteFile("~output.json", rankingsJson, 0644)

}
