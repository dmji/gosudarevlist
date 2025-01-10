package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *service) UpdateItemsFromCategory(ctx context.Context, category enums.Category, mode model.CategoryUpdateMode) error {
	logger.Infow(ctx, "Update Target Category | Pipe started", "category", category, "mode", mode)

loop_pages:
	for iPage := 1; ; iPage++ {
		items, err := s.animelayerApi.GetItemsFromCategoryPages(ctx, category, iPage)
		if err != nil {
			logger.Infow(ctx, "Update Target Category | Items getting failed", "category", category, "page", iPage, "error", err)
			return err
		}

		if len(items) == 0 {
			break
		}

	loop_items:
		for _, item := range items {

			_, err = s.repo.GetItemByIdentifier(ctx, item.Identifier)
			bInsert := errors.Is(err, sql.ErrNoRows)
			if bInsert {
				err = s.repo.InsertItem(ctx, item, category)
			} else {
				err = s.repo.UpdateItem(ctx, item)
			}

			if _, ok := repository.IsErrorItemNotChanged(err); ok {
				switch mode {

				case model.CategoryUpdateModeAll:
					continue loop_items
				case model.CategoryUpdateModeWhileNew:
					fallthrough
				default:
					break loop_pages
				}
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
	}

	logger.Infow(ctx, "Update Target Category | Pipe finished without errors", "category", category, "mode", mode)
	return nil
}