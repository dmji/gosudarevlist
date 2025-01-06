package middleware

import (
	"log"
	"maps"
	"net/http"
	"net/url"
	"slices"

	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func HxPushUrlMiddleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		currentUri := r.Header.Get("HX-Current-URL")
		currentUrl, err := url.Parse(currentUri)
		if err != nil {
			log.Panic(err)
		}
		logger.Infow(r.Context(), "queries", "current", currentUrl.Query().Encode(), "request", r.URL.Query().Encode())
		currentQuery := currentUrl.Query()
		for key, values := range r.URL.Query() {
			currentQuery[key] = values
		}
		for key := range currentQuery {
			currentQuery[key] = slices.DeleteFunc(currentQuery[key], func(value string) bool { return len(value) == 0 })
		}
		maps.DeleteFunc(currentQuery, func(key string, value []string) bool { return len(value) == 0 })

		log.Printf("Middleware Hx-Push-Url | Query URI: %v", currentQuery.Encode())
		newUri := currentUrl.Path + custom_url.QueryOrEmpty(currentQuery.Encode())

		w.Header().Set("Access-Control-Expose-Headers", "Hx-Push-Url")
		w.Header().Set("Hx-Push-Url", newUri)

		log.Printf("Middleware Hx-Push-Url | Prev URI: %s", currentUri)
		log.Printf("Middleware Hx-Push-Url | New URI: %s", newUri)

		r.URL.RawQuery = currentQuery.Encode()
		handler(w, r)
	}
}
