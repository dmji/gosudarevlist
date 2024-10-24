package filters

import (
	"collector/pkg/custom_url"
	"collector/pkg/logger"
	"context"
	"net/url"
	"reflect"
	"strconv"

	"github.com/google/go-querystring/query"
)

type ApiCardsParams struct {
	Page        int    `url:"page"`
	SearchQuery string `url:"query"`
}

func ParseApiCardsParams(ctx context.Context, q url.Values, defaultPage int) ApiCardsParams {
	res := ApiCardsParams{}

	q = custom_url.QueryCustomParse(q)

	if s := q.Get(res.getUrlTagByFieldName("Page")); len(s) > 0 {
		page, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			logger.Errorw(ctx, "parsing page string to int", "error", err)
		}
		res.Page = int(page)
	} else {
		res.Page = defaultPage
	}

	if s := q.Get(res.getUrlTagByFieldName("SearchQuery")); len(s) > 0 {
		res.SearchQuery = s
	}

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
