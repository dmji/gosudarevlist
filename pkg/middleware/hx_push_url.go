package middleware

import (
	"log"
	"net/http"
	"net/url"

	"github.com/dmji/gosudarevlist/pkg/custom_url"
)

func HxPushUrlMiddleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		currentUri := r.Header.Get("HX-Current-URL")
		currentUrl, err := url.Parse(currentUri)
		if err != nil {
			log.Panic(err)
		}

		mergedQuery := custom_url.QueryCustomParse(r.URL.Query())

		log.Printf("Middleware Hx-Push-Url | Query URI: %v", mergedQuery.Encode())
		newUri := currentUrl.Path + custom_url.QueryValuesToString(&mergedQuery)

		w.Header().Set("Access-Control-Expose-Headers", "Hx-Push-Url")
		w.Header().Set("Hx-Push-Url", newUri)

		log.Printf("Middleware Hx-Push-Url | Prev URI: %s", currentUri)
		log.Printf("Middleware Hx-Push-Url | New URI: %s", newUri)

		handler(w, r)
	}
}
