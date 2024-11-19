package query_updates

import (
	"collector/internal/custom_types"
	"collector/pkg/custom_url"
	"collector/pkg/logger"
	"context"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

type ApiUpdateParams struct {
	Page        custom_types.Page `url:"page"`
	SearchQuery string            `url:"query"`
}

func (p ApiUpdateParams) IsCompletedUrl() string {
	return p.getUrlTagByFieldName("IsCompleted")
}

func Parse(ctx context.Context, q url.Values, defaultPage int) *ApiUpdateParams {
	res := &ApiUpdateParams{}

	q = custom_url.QueryCustomParse(q)

	err := res.Page.DecodeValues(q.Get(res.getUrlTagByFieldName("Page")), defaultPage)
	if err != nil {
		logger.Errorw(ctx, "parsing page string to int", "error", err)
	}

	res.SearchQuery = q.Get(res.getUrlTagByFieldName("SearchQuery"))

	return res
}

func (p *ApiUpdateParams) Values(ctx context.Context) url.Values {

	qv, err := query.Values(p)
	if err != nil {
		logger.Errorw(ctx, "encode values to query", "error", err)
		return qv
	}

	return custom_url.QueryCustomParse(qv)
}

func (p ApiUpdateParams) getUrlTagByFieldName(fieldName string) string {
	f, ok := reflect.TypeOf(&p).Elem().FieldByName(fieldName)
	if !ok {
		return ""
	}

	return f.Tag.Get("url")
}
