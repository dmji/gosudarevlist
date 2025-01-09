//go:build !embed

package lang

import "context"

func (s *Storage) Reload(ctx context.Context) {
	s.bundle.MustLoadMessageFile("lang/translations/active.ru.yaml")
}
