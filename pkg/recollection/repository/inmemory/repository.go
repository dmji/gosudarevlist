package repository_inmemory

import (
	animelayer_model "collector/pkg/animelayer/model"
	_ "embed"
	"encoding/json"
)

//go:embed db/items.json
var content []byte

//go:embed db/descriptions.json
var descriptions []byte

type repository struct {
	db           []animelayer_model.Item
	descriptions []animelayer_model.Description
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
