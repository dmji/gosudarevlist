package expose_header_utils

import (
	"context"
	"net/http"
	"net/url"

	"github.com/dmji/gosudarevlist/pkg/logger"
)

func HxReplaceUrl(ctx context.Context, w http.ResponseWriter, r *http.Request, fnOverrideRawQuery func(q string) (string, error)) (*url.URL, error) {
	currentUri := r.Header.Get("HX-Current-URL")
	currentUrl, err := url.Parse(currentUri)
	if err != nil {
		logger.Errorw(ctx, "Hx-Replace-Url | Url parse failed", "error", err)
		return nil, err
	}

	currentUrl.RawQuery, err = fnOverrideRawQuery(currentUrl.RawQuery)
	if err != nil {
		logger.Errorw(ctx, "Hx-Replace-Url | Url encode failed", "error", err)
		return nil, err
	}

	s := currentUrl.String()
	WriterExposeHeader(w, "Hx-Replace-Url", s)
	logger.Infow(ctx, "Hx-Replace-Url | Pushed Url", "from", currentUri, "to", s)

	return currentUrl, nil
}
