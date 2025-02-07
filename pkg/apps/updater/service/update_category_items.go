package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dmji/go-animelayer-parser"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *service) UpdateItemsFromCategory(ctx context.Context, category enums.Category, mode model.CategoryUpdateMode) error {
	data := s.updaterDataByCategory(ctx, category)
	bOk := data.mx.TryLock()
	if !bOk {
		return NewRrrorInProcess(category, data.lastUpdateTimer)
	}
	defer data.mx.Unlock()

	data.lastUpdateTimer = time.Now()
	_ = data.category
	go s.managerNotifier.UpdateTrigger(ctx, category)
	logger.Infow(ctx, "Update Target Category | Pipe started", "category", category, "mode", mode)

	updatedIdentifiers := make(map[string]interface{})

loop_pages:
	for iPage := 1; ; iPage++ {
		items, err := s.animelayerApi.GetItemsFromCategoryPages(ctx, category, iPage)
		if errors.Is(err, animelayer.ErrorEmptyPage) {
			break loop_pages
		}

		if err != nil {
			logger.Infow(ctx, "Update Target Category | Items getting failed", "category", category, "page", iPage, "error", err)
			return err
		}

		if len(items) == 0 {
			break
		}
		logger.Infow(ctx, "Update Target Category | Pipe in-progress", "category", category, "mode", mode, "page", iPage)

		nItemWasSame := 0
		for i := range items {
			item := items[len(items)-i-1]

			if _, ok := updatedIdentifiers[item.Identifier]; ok {
				break loop_pages
			}
			updatedIdentifiers[item.Identifier] = nil

			_, err = s.repo.GetItemByIdentifier(ctx, item.Identifier)
			bInsert := errors.Is(err, sql.ErrNoRows)
			if bInsert {
				err = s.repo.InsertItem(ctx, item, category)
			} else {
				err = s.repo.UpdateItem(ctx, item)
			}

			if _, ok := repository.IsErrorItemNotChanged(err); ok {
				nItemWasSame++
				continue
			}

			if err != nil {
				if bInsert {
					logger.Infow(ctx, "Update Target Category | Item insertion failed", "identifier", item.Identifier, "error", err)
				} else {
					logger.Infow(ctx, "Update Target Category | Item updating failed", "identifier", item.Identifier, "error", err)
				}
				return err
			}

			if bInsert {
				logger.Infow(ctx, "Update Target Category | Item inserted as new", "identifier", item.Identifier)
			} else {
				logger.Infow(ctx, "Update Target Category | Item has being updated", "identifier", item.Identifier)
			}
		}

		if nItemWasSame == len(items) {
			switch mode {
			case model.CategoryUpdateModeAll:
				continue
			case model.CategoryUpdateModeWhileNew:
				fallthrough
			default:
				break loop_pages
			}
		}
	}

	logger.Infow(ctx, "Update Target Category | Pipe finished without errors", "category", category, "mode", mode)
	return nil
}
