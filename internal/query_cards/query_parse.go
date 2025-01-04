package query_cards

import (
	"context"
	"net/url"
	"strings"

	"github.com/dmji/gosudarevlist/internal/custom_types"
	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"

	"github.com/google/go-querystring/query"
)

func Parse(ctx context.Context, q url.Values, defaultPage custom_types.Page) *ApiCardsParams {
	res := &ApiCardsParams{
		Page: defaultPage,
	}

	q = custom_url.QueryCustomParse(q)

	err := res.Page.UnmarshalText([]byte(q.Get(res.getUrlTagByFieldName("Page"))))
	if err != nil {
		logger.Errorw(ctx, "parsing page string to int", "error", err)
	}

	res.SearchQuery = q.Get(res.getUrlTagByFieldName("SearchQuery"))

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

	if statusesList := q.Get(res.getUrlTagByFieldName("Statuses")); len(statusesList) > 0 {
		if statuses := strings.Split(statusesList, " "); len(statuses) != 0 {
			for _, status := range statuses {
				c, err := model.ReleaseStatusFromString(status)
				if err != nil {
					logger.Errorw(ctx, "failed parse query status", "error", err)
					continue
				}

				res.Statuses = append(res.Statuses, c)
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
