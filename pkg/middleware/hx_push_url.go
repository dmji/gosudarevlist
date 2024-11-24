package middleware

import (
	"collector/pkg/custom_url"
	"log"
	"net/http"
	"net/url"
)

func FuncHxPushUrl(currentUri string, queryStr url.Values, w http.ResponseWriter, fnModify ...func(v *url.Values)) string {
	currentUrl, err := url.Parse(currentUri)
	if err != nil {
		log.Panic(err)
	}

	mergedQuery := custom_url.QueryCustomParse(queryStr)
	for _, fn := range fnModify {
		fn(&mergedQuery)
	}

	log.Printf("Middleware Hx-Push-Url | Query URI: %v", mergedQuery.Encode())
	newUri := currentUrl.Path + custom_url.QueryValuesToString(&mergedQuery)

	w.Header().Set("Access-Control-Expose-Headers", "Hx-Push-Url")
	w.Header().Set("Hx-Push-Url", newUri)

	return newUri
}

func HxPushUrlMiddleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		currentUri := r.Header.Get("HX-Current-URL")
		log.Printf("Middleware Hx-Push-Url | Prev URI: %s", currentUri)

		newUri := FuncHxPushUrl(currentUri, r.URL.Query(), w)
		log.Printf("Middleware Hx-Push-Url | New URI: %s", newUri)

		handler(w, r)
	}
}
