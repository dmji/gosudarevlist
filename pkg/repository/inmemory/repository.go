package repository_inmemory

import (
	"collector/pkg/model"
	_ "embed"
	"encoding/json"
)

//go:embed db/test.json
var content []byte

//go:embed db/descriptions.json
var descriptions []byte

type repository struct {
	db           []model.AnimeLayerItem
	descriptions []model.AnimeLayerItemDescription
}

func New() *repository {
	res := &repository{}
	err := json.Unmarshal(content, &res.db)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(descriptions, &res.descriptions)
	if err != nil {
		panic(err)
	}

	return res
}
