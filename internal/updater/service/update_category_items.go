package service

import (
	"context"
	"errors"

	"github.com/dmji/go-animelayer-parser"
	"github.com/dmji/gosudarevlist/internal/updater/model"
	"github.com/dmji/gosudarevlist/internal/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *service) UpdateItems(ctx context.Context, mode model.CategoryUpdateMode, nPageLimit int) error {
	err := s.checkMx()
	if err != nil {
		return err
	}
	defer s.mx.Unlock()

	logger.Infow(ctx, "Update Target Category | Pipe started", "category", s.category, "mode", mode)

	updatedIdentifiers := make(map[string]interface{})

loop_pages:
	for iPage := 1; ; iPage++ {
		if nPageLimit != 0 && iPage > nPageLimit {
			break
		}

		items, err := s.animelayerApi.GetItemsFromCategoryPages(ctx, s.category, iPage)
		if errors.Is(err, animelayer.ErrorEmptyPage) {
			break loop_pages
		}

		if err != nil {
			logger.Infow(ctx, "Update Target Category | Items getting failed", "category", s.category, "page", iPage, "error", err)
			return err
		}

		if len(items) == 0 {
			break
		}
		logger.Infow(ctx, "Update Target Category | Pipe in-progress", "category", s.category, "mode", mode, "page", iPage)

		nItemWasSame := 0
		for i := range items {
			item := items[len(items)-i-1]

			// Get new item data from AnimeLayer
			if _, ok := updatedIdentifiers[item.Identifier]; ok {
				break loop_pages
			}
			updatedIdentifiers[item.Identifier] = nil

			// Get current item data from DB
			_, err = s.repo.GetItemByIdentifier(ctx, item.Identifier)
			method := getMethod(err)
			err = updateItem(ctx, method, s.repo, item, s.category)

			if _, ok := repository.IsErrorItemNotChanged(err); ok {
				nItemWasSame++
				continue
			}

			if err != nil {
				logger.Errorw(ctx, "update target item | item processing failed",
					"method", method.String(),
					"identifier", item.Identifier,
					"error", err,
				)
				return err
			}

			logger.Infow(ctx, "update target item | item processing finished",
				"method", method.String(),
				"identifier", item.Identifier,
			)
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

	logger.Infow(ctx, "Update Target Category | Pipe finished without errors", "category", s.category, "mode", mode)
	return nil
}
