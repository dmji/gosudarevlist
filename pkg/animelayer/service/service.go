package animelayer_service

import (
	"collector/internal/animelayer_parser"
	animelayer_model "collector/pkg/animelayer/model"
	"collector/pkg/parser"
	"context"
	"fmt"
	"net/http"
	"os"

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

func getTestCreadentials() parser.Credentials {
	return parser.Credentials{
		Login:    os.Getenv("loginAnimeLayer"),
		Password: os.Getenv("passwordAnimeLayer"),
	}
}

func New() *service {
	client, err := parser.HttpClientWithAuth(
		BaseUrl,
		getTestCreadentials(),
	)
	if err != nil {
		panic(err)
	}

	return &service{
		client: *client,
	}
}

type service struct {
	client http.Client
}

func (s *service) GetItems(ctx context.Context, category string) ([]animelayer_model.Description, error) {
	return nil, nil
}

func (s *service) GetDescription(ctx context.Context, guid string) (*animelayer_model.Description, error) {

	url := FormatUrlToItem(guid)

	doc, err := parser.LoadHtmlDocument(&s.client, url)
	if err != nil {
		return nil, err
	}

	item := animelayer_parser.ParseItem(ctx, doc)
	if item == nil {
		return nil, fmt.Errorf("error on parse document of guid='%s'", guid)
	}

	return item, nil
}

func (s *service) StartCategoryParsing(ctx context.Context, category string) (*animelayer_model.Description, error) {
	return nil, nil
}
