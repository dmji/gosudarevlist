package repository

import (
	"collector/pkg/model"
	"context"
	"log"
)

func (r *repository) GetItems(ctx context.Context, count int, offset int) ([]model.AnimeLayerItem, error) {
	log.Printf("In-Memory repo | GetItems count: %d, offset: %d, db items: %d", count, offset, len(r.db))

	if rest := len(r.db) - offset; rest < count {
		count = rest
	}

	res := r.db[offset : offset+count]
	log.Printf("In-Memory repo | GetItems result items: %d", len(res))
	return res, nil
}
