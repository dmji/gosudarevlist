package repository

import (
	"collector/pkg/model"
	"context"
)

func (r *repository) GetItems(ctx context.Context, count int, offset int) ([]model.AnimeLayerItem, error) {

	if rest := len(r.db) - offset; rest < count {
		count = rest
	}

	res := r.db[offset : offset+count]
	return res, nil
}
