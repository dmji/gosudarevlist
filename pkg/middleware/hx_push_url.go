package middleware

import (
	"maps"
	"net/http"
	"net/url"
	"slices"

	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func HxPushUrlMiddleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		currentUri := r.Header.Get("HX-Current-URL")
		currentUrl, err := url.Parse(currentUri)
		if err != nil {
			logger.Errorw(ctx, "Middleware Hx-Push-Url | Url Parse failed", "error", err)
		}

		currentQuery, err := custom_url.QueryCustomParse(currentUrl.RawQuery)
		if err != nil {
			logger.Errorw(ctx, "Middleware Hx-Push-Url | Query Parse failed", "error", err)
		}

		// move parameters from request
		for key, values := range r.URL.Query() {
			currentQuery[key] = values
		}
		// remove empty values
		for key := range currentQuery {
			currentQuery[key] = slices.DeleteFunc(currentQuery[key], func(value string) bool { return len(value) == 0 })
		}
		maps.DeleteFunc(currentQuery, func(key string, value []string) bool { return len(value) == 0 })

		newQueryStr := custom_url.QueryCustomEncode(currentQuery)
		newUri := currentUrl.Path + custom_url.QueryOrEmpty(newQueryStr)

		w.Header().Set("Access-Control-Expose-Headers", "Hx-Push-Url")
		w.Header().Set("Hx-Push-Url", newUri)
		logger.Infow(ctx, "Middleware Hx-Push-Url | Pushed Url", "from", currentUri, "to", newUri)

		r.URL.RawQuery = newQueryStr
		handler(w, r)
	}
}
