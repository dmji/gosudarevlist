package custom_url

import (
	"context"
	"maps"
	"net/url"
	"slices"

	"github.com/dmji/gosudarevlist/pkg/logger"
)

func MergeQueryStringWithExtraQuery(ctx context.Context, currentQueryStr string, extraQueryParams url.Values) string {
	currentQuery, err := QueryCustomParse(currentQueryStr)
	if err != nil {
		logger.Errorw(ctx, "Middleware Hx-Replace-Url | Query Parse failed", "error", err)
	}

	// move parameters from request
	for key, values := range extraQueryParams {
		currentQuery[key] = values
	}

	// remove empty values
	for key := range currentQuery {
		currentQuery[key] = slices.DeleteFunc(currentQuery[key], func(value string) bool { return len(value) == 0 })
	}
	maps.DeleteFunc(currentQuery, func(key string, value []string) bool { return len(value) == 0 })

	return QueryCustomEncode(currentQuery)
}
