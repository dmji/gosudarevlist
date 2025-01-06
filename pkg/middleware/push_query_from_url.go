package middleware

import (
	"net/http"
	"net/url"

	"github.com/dmji/gosudarevlist/pkg/logger"
)

func PushQueryFromUrlMiddleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		currentUri := r.Header.Get("HX-Current-URL")
		currentUrl, err := url.Parse(currentUri)
		if err != nil {
			logger.Errorw(ctx, "Middleware Hx-Push-Url | Url Parse failed", "error", err)
		}

		r.URL.RawQuery = currentUrl.RawQuery
		handler(w, r)
	}
}
