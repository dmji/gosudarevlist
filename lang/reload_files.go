//go:build !embed

package lang

func (s *Storage) Reload() {
	s.bundle.MustLoadMessageFile("lang/active.ru.yaml")
}
