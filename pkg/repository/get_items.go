package repository

import (
	"collector/pkg/model"
	"context"
	"log"
)

func (r *repository) GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.AnimeLayerItem, error) {
	log.Printf("In-Memory repo | GetItems count: %d, offset: %d, db items: %d", opt.Count, opt.Offset, len(r.db))

	if rest := len(r.db) - opt.Offset; rest < opt.Count {
		opt.Count = rest
	}

	res := r.db[opt.Offset : opt.Offset+opt.Count]
	log.Printf("In-Memory repo | GetItems result items: %d", len(res))
	return res, nil

	/* log.Printf("In-Memory repo | GetItems count: %d, offset: %d, db items: %d", opt.Count, opt.Offset, len(r.db))

	if opt.Count == 0 {
		return nil, nil
	}

	res := make([]model.AnimeLayerItem, 0, opt.Count)
	currentOffset := 0
	for i := 0; i < len(r.db); i++ {

		item := r.db[i]

		if len(opt.SearchQuery) > 0 {

			if dist := fuzzy.RankMatch(opt.SearchQuery, item.Name); dist == -1 || dist < int(len(opt.SearchQuery)*75/100) {
				continue
			}

		}

		if currentOffset < opt.Offset {
			log.Printf("In-Memory repo | GetItems offset shift from '%d' to '%d'", currentOffset, opt.Offset)
			currentOffset++
			continue
		}

		res = append(res, item)
		if len(res) >= opt.Count {
			break
		}

	}

	log.Printf("In-Memory repo | GetItems result items: %d", len(res))
	return res, nil */
}
