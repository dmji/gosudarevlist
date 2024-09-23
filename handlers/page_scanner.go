package handlers

import (
	"collector/internal/animelayer_parser"
	animelayer_model "collector/pkg/animelayer/model"
	animelayer_service "collector/pkg/animelayer/service"
	"collector/pkg/parser"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

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

func (s *router) ScannerPageHandler(w http.ResponseWriter, r *http.Request) {
	//
	log.Print("Handler Scanner | Reached")
	items := make([]animelayer_model.Item, 0, 2000)

	mx := sync.Mutex{}

	wg, ctx := errgroup.WithContext(r.Context())
	wg.SetLimit(20)

	wg.Go(func() error {
		log.Print("Main loop started")

		for i := range 40 {
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
							url := animelayer_service.FormatUrlToItemsPage("/torrents/anime", i+1)
							log.Print(url)

							client, err := parser.HttpClientWithAuth(
								animelayer_service.BaseUrl,
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
							pageItems := animelayer_parser.ParseCategory(ctx, doc)
							if len(pageItems) == 0 {
								log.Printf("Reading %d page finished with error", i+1)
								return errors.New("end of context")
							}

							mx.Lock()
							defer mx.Unlock()
							items = append(items, pageItems...)
							log.Printf("Reading %d page finished", i+1)
							return nil
						}
					})
			}
		}
		return nil
	})
	_ = wg.Wait()
	log.Printf("Handler Scanner | Items %d", len(items))

	//
	/* 	repoItems, _ := s.s.Repo().GetItems(ctx, model.OptionsGetItems{Count: 10000, Offset: 0})
	   	result := make([]pages.ScanResult, 0, len(items))
	   	for i, item := range items {
	   		status := ""
	   		if item.Completed {
	   			status = "Completed"
	   		} else {
	   			status = "In Progress"
	   		}

	   		repoIndex := slices.IndexFunc(repoItems, func(ri animelayer_model.Item) bool { return ri.GUID == item.GUID })
	   		log.Printf("Handler Scanner | Items with guid='%s' found in repo '%d'", item.GUID, repoIndex)
	   		if repoIndex == -1 {
	   			log.Printf("Handler Scanner | Items with guid='%s' skipped", item.GUID)
	   			continue
	   		}
	   		if /* repoIndex != -1 &&  */ /* repoItems[repoIndex].Name == item.Name && repoItems[repoIndex].Completed == item.Completed { */
	/* 	log.Printf("Handler Scanner | Items with guid='%s' skipped", item.GUID)
			continue
		}
		result = append(result, pages.ScanResult{
			ID:     i,
			GUID:   item.GUID,
			Status: status,
			Title:  item.Name,
		})
	}
	log.Printf("Handler Scanner | Items updated %d", len(result)) */

	/* 	err := pages.Scanner(result).Render(r.Context(), w)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

}
