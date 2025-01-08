package middleware

import (
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/expose_header_utils"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func LangerToContextMiddleware(storage *lang.Storage) func(http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			langStr := expose_header_utils.GetCookiePreferedLanguage(w, r).Value
			l, err := lang.TagLangFromString(langStr)
			if err != nil {
				logger.Errorw(r.Context(), "Middleware Langer To Context | Lang tag parsing failed", "string", langStr, "error", err)
				l = lang.TagEnglish
			}
			ctx := storage.ToContext(r.Context(), l)

			handler(w, r.WithContext(ctx))
		}
	}
}
