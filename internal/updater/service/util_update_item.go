package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dmji/gosudarevlist/internal/updater/model"
	"github.com/dmji/gosudarevlist/internal/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

func updateItem(ctx context.Context, method enums.UpdateMethod, repo repository.AnimeLayerRepositoryDriver, item *model.AnimelayerItem, category enums.Category) error {
	switch method {
	case enums.UpdateMethodInsertion:
		return repo.InsertItem(ctx, item, category)
	case enums.UpdateMethodUpdating:
		return repo.UpdateItem(ctx, item)
	default:
		panic(fmt.Sprintf("unexpected enums.UpdateMethod: %#v", method))
	}
}

func getMethod(err error) enums.UpdateMethod {
	if errors.Is(err, sql.ErrNoRows) {
		return enums.UpdateMethodInsertion
	}

	return enums.UpdateMethodUpdating
}
