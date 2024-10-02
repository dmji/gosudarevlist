package repository_inmemory

import (
	_ "embed"
	"encoding/json"

	"github.com/dmji/go-animelayer-parser"
)

//go:embed db/items.json
var content []byte

//go:embed db/descriptions.json
var descriptions []byte

type repository struct {
	db           []animelayer.ItemPartial
	descriptions []animelayer.ItemDetailed
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
