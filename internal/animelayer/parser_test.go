package animelayer_test

import (
	animelayer_pages_parser "collector/internal/animelayer/pages_parser"
	"collector/pkg/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"golang.org/x/sync/errgroup"
)

const (
	BASE_ADDRESS_URI = "https://animelayer.ru"
)

func TestMainTest(t *testing.T) {

	items := make([]model.AnimeLayerItem, 0, 2000)
	ctx := context.Background()

	mx := sync.Mutex{}

	i := 0
	wg, ctx := errgroup.WithContext(ctx)
	wg.SetLimit(20)

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

							pageItems := animelayer_pages_parser.CollectBaseItemsFromAddress(ctx, BASE_ADDRESS_URI, "anime", i+1)
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

	err := wg.Wait()
	fmt.Printf("Errgroup error: %v", err)
	rankingsJson, _ := json.Marshal(items)
	_ = os.WriteFile("~output.json", rankingsJson, 0644)

}
