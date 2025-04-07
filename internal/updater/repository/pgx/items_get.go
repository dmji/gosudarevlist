package repository_pgx

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dmji/gosudarevlist/internal/updater/model"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"
)

func (r *repository) GetItems(ctx context.Context, from time.Time) ([]*model.AnimelayerItem, error) {
	pgTime, err := pgx_utils.TimeToPgTimestamp(&from)
	if err != nil {
		return nil, err
	}

	items, err := r.query.GetItems(ctx, pgTime)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItemByIdentifier", "error", err)
		return nil, err
	}

	res := make([]*model.AnimelayerItem, 0, len(items))
	for i := range items {
		res = append(res, pgItemFromItem(ctx, items[i]))
	}

	return res, nil
}
