//go:build !embed

package assets

import "net/http"

func Handler() http.Handler {
	return http.FileServer(http.Dir("assets"))
}
