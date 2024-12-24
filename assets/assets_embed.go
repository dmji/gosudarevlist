//go:build embed

package assets

import (
	"embed"
	_ "embed"
	"net/http"
)

//go:embed css/*
//go:embed images/*
var content embed.FS

func Handler() http.Handler {
	return http.FileServerFS(content)
}
