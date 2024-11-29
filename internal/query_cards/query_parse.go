package query_cards

import (
	"collector/pkg/custom_url"
	"collector/pkg/logger"
	"collector/pkg/recollection/model"
	"context"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

func Parse(ctx context.Context, q url.Values, defaultPage int) *ApiCardsParams {
	res := &ApiCardsParams{}

	q = custom_url.QueryCustomParse(q)

	err := res.Page.DecodeValues(q.Get(res.getUrlTagByFieldName("Page")), defaultPage)
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
				c, err := model.StatusFromString(status)
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
