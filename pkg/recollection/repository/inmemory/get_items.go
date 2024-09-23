package repository_inmemory

import (
	animelayer_model "collector/pkg/animelayer/model"
	"collector/pkg/recollection/model"
	"context"
	"log"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func (r *repository) GetItems(ctx context.Context, opt model.OptionsGetItems) ([]animelayer_model.Item, error) {
	log.Printf("In-Memory repo | GetItems count: %d, offset: %d, db items: %d", opt.Count, opt.Offset, len(r.db))

	if opt.Count == 0 {
		return nil, nil
	}

	res := make([]animelayer_model.Item, 0, opt.Count)
	currentOffset := 0
	for i := 0; i < len(r.db); i++ {

		item := r.db[i]

		if len(opt.SearchQuery) > 0 {
			if !fuzzy.MatchNormalizedFold(opt.SearchQuery, item.Name) {
				continue
			}
		}

		if currentOffset < opt.Offset {
			currentOffset++
			continue
		}

		res = append(res, item)
		if len(res) >= opt.Count {
			break
		}

	}

	log.Printf("In-Memory repo | GetItems result items: %d", len(res))
	return res, nil
}
