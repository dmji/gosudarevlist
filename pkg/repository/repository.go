package repository

import (
	"collector/pkg/model"
	"context"
	_ "embed"
	"encoding/json"
)

//go:embed db/test.json
var content []byte

type repository struct {
	db []model.AnimeLayerItem
}

func New() *repository {
	res := &repository{}
	err := json.Unmarshal(content, &res.db)

	if err != nil {
		panic(err)
	}

	return res
}

type AnimeLayerRepositoryDriver interface {
	GetItems(ctx context.Context, count int, offset int) ([]model.AnimeLayerItem, error)
}
