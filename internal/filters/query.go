package filters

import (
	"collector/internal/custom_types"
	"collector/pkg/custom_url"
	"collector/pkg/logger"
	"context"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

type ApiCardsParams struct {
	Page        custom_types.Page        `url:"page"`
	SearchQuery string                   `url:"query"`
	IsCompleted *custom_types.BoolExProp `url:"status"`
}

func ParseApiCardsParams(ctx context.Context, q url.Values, defaultPage int) *ApiCardsParams {
	res := &ApiCardsParams{
		IsCompleted: &custom_types.BoolExProp{
			Name: "completed",
		},
	}

	q = custom_url.QueryCustomParse(q)

	err := res.Page.DecodeValues(q.Get(res.getUrlTagByFieldName("Page")), defaultPage)
	if err != nil {
		logger.Errorw(ctx, "parsing page string to int", "error", err)
	}

	res.SearchQuery = q.Get(res.getUrlTagByFieldName("SearchQuery"))
	res.IsCompleted.DecodeValues(q.Get(res.getUrlTagByFieldName("IsCompleted")))

	return res
}

func (p *ApiCardsParams) Values(ctx context.Context) url.Values {

	qv, err := query.Values(p)
	if err != nil {
		logger.Errorw(ctx, "encode values to query", "error", err)
		return qv
	}

	return custom_url.QueryCustomParse(qv)
}

func (p ApiCardsParams) getUrlTagByFieldName(fieldName string) string {
	f, ok := reflect.TypeOf(&p).Elem().FieldByName(fieldName)
	if !ok {
		return ""
	}

	return f.Tag.Get("url")
}
