//go:build embed

package lang

import (
	"context"
	"embed"

	"github.com/dmji/gosudarevlist/pkg/logger"
)

//go:embed translations/*
var content embed.FS

func (s *Storage) Reload(ctx context.Context) {
	_, err := s.bundle.LoadMessageFileFS(content, "translations/active.ru.yaml")
	if err != nil {
		logger.Errorw(ctx, "String Multilang Storage refresh failed", "path", "translations/active.ru.yaml", "err", err)
	}
}
