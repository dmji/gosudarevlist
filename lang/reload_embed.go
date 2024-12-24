//go:build embed

package lang

import (
	"embed"
)

//go:embed translations/*
var content embed.FS

func (s *Storage) Reload() {
	s.bundle.LoadMessageFileFS(content, "active.ru.yaml")
}
