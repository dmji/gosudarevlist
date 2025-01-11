package repository_pgx

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"
)

func (r *repository) GetLastCategoryUpdateItem(ctx context.Context, category enums.Category) (*time.Time, error) {
	ts, err := r.query.GetLastCategoryUpdateItem(ctx, categoriesToAnimelayerCategories([]enums.Category{category}, false)[0])
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItemByIdentifier", "error", err)
		return nil, err
	}

	return pgx_utils.TimeFromPgTimestamp(ts), nil
}
