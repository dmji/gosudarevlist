package query_cards

import (
	"collector/internal/custom_types"
	"collector/pkg/custom_url"
	"collector/pkg/logger"
	"collector/pkg/recollection/model"
	"context"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

type ApiCardsParams struct {
	Page        custom_types.Page        `url:"page"`
	SearchQuery string                   `url:"query"`
	IsCompleted *custom_types.BoolExProp `url:"status"`
	ShowFilters bool                     `url:"show_filters,omitempty"`
	Categories  []model.Category         `url:"category,space"`
}

func (p ApiCardsParams) IsCompletedUrl() string {
	return p.getUrlTagByFieldName("IsCompleted")
}

func (p ApiCardsParams) CategoriesUrl() string {
	return p.getUrlTagByFieldName("Categories")
}

func Parse(ctx context.Context, q url.Values, defaultPage int) *ApiCardsParams {
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

	if categoriesList := q.Get(res.getUrlTagByFieldName("Categories")); len(categoriesList) > 0 {
		if categories := strings.Split(categoriesList, " "); len(categories) != 0 {
			for _, category := range categories {
				c, err := model.CategoryFromString(category)
				if err != nil {
					logger.Errorw(ctx, "failed parse query category", "error", err)
					continue
				}

				res.Categories = append(res.Categories, c)
			}
		}
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

	rs := f.Tag.Get("url")
	n, _, _ := strings.Cut(rs, ",")
	return n
}
