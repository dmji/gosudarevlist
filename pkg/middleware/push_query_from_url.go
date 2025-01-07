package middleware

import (
	"net/http"
	"net/url"

	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func PushQueryFromUrlMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		currentUri := r.Header.Get("HX-Current-URL")
		currentUrl, err := url.Parse(currentUri)
		if err != nil {
			logger.Errorw(ctx, "Middleware Hx-Push-Url | Url Parse failed", "error", err)
		}

		r.URL.RawQuery = custom_url.MergeQueryStringWithExtraQuery(ctx, currentUrl.RawQuery, r.URL.Query())
		handler(w, r)
	}
}
