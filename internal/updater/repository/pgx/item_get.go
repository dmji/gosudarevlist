package repository_pgx

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dmji/gosudarevlist/internal/updater/model"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (r *repository) GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error) {
	item, err := r.query.GetItemByIdentifier(ctx, identifier)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItemByIdentifier", "error", err)
		return nil, err
	}

	return pgItemFromItem(ctx, item), nil
}
