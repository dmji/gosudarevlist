package service

import (
	"context"
	"fmt"

	"github.com/dmji/gosudarevlist/internal/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *service) UpdateTargetItem(ctx context.Context, identifier string) error {
	err := s.checkMx()
	if err != nil {
		return err
	}
	defer s.mx.Unlock()

	// Get new item data from AnimeLayer
	item, err := s.animelayerApi.GetItemByIdentifier(ctx, identifier)
	if err != nil {
		logger.Infow(ctx, "update target item | item getting failed", "category", s.category, "identifier", identifier, "error", err)
		return err
	}
	if item.Category != s.category {
		logger.Warnw(ctx, "update target item | item category not match updater")
		return fmt.Errorf("item category='%s', but expected for updater is '%d'", item.Category, s.category)
	}

	// Get current item data from DB
	_, err = s.repo.GetItemByIdentifier(ctx, identifier)
	method := getMethod(err)
	err = updateItem(ctx, method, s.repo, item, s.category)

	if _, ok := repository.IsErrorItemNotChanged(err); ok {
		return nil
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

	return nil
}
