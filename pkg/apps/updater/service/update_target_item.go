package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *service) UpdateTargetItem(ctx context.Context, identifier string, category enums.Category) error {
	data := s.updaterDataByCategory(ctx, category)
	bOk := data.mx.TryLock()
	if !bOk {
		return NewRrrorInProcess(category, data.lastUpdateTimer)
	}
	defer data.mx.Unlock()

	item, err := s.animelayerApi.GetItemByIdentifier(ctx, identifier)
	if err != nil {
		logger.Infow(ctx, "Update Target Item | Item getting failed", "category", category, "identifier", identifier, "error", err)
		return err
	}

	_, err = s.repo.GetItemByIdentifier(ctx, identifier)
	bInsert := errors.Is(err, sql.ErrNoRows)
	if bInsert {
		err = s.repo.InsertItem(ctx, item, category)
	} else {
		err = s.repo.UpdateItem(ctx, item)
	}

	if err != nil {
		if bInsert {
			logger.Infow(ctx, "Update Target Item | Item insertion failed", "identifier", item.Identifier, "error", err)
		} else {
			logger.Infow(ctx, "Update Target Item | Item updating failed", "identifier", item.Identifier, "error", err)
		}
		return err
	}

	if bInsert {
		logger.Infow(ctx, "Update Target Item | Item inserted as new", "identifier", identifier)
	} else {
		logger.Infow(ctx, "Update Target Item | Item has being updated", "identifier", identifier)
	}

	return nil
}
