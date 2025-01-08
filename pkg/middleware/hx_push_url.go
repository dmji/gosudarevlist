package middleware

import (
	"net/http"

	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/expose_header_utils"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func HxReplaceUrlMiddleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		newUrl, err := expose_header_utils.HxReplaceUrl(ctx, w, r, func(q string) (string, error) {
			return custom_url.MergeQueryStringWithExtraQuery(ctx, q, r.URL.Query()), nil
		})
		if err != nil {
			logger.Errorw(ctx, "HxPushUrl failed", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Infow(ctx, "Middleware Hx-Replace-Url | Pushed Url", "from", r.Header.Get("HX-Current-URL"), "to", newUrl.String())
		handler(w, r)
	}
}
