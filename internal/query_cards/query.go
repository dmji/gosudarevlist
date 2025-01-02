package query_cards

import (
	"context"
	"reflect"
	"strings"

	"github.com/dmji/gosudarevlist/internal/custom_types"
	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ApiCardsParams struct {
	Page        custom_types.Page `url:"page"`
	SearchQuery string            `url:"query"`

	Categories []model.Category      `url:"category,space"`
	Statuses   []model.ReleaseStatus `url:"release_status,space"`
}

func StatusesUrl() string {
	return ApiCardsParams{}.getUrlTagByFieldName("Statuses")
}

func StatusesPresentation(ctx context.Context) string {
	return lang.Message(ctx, &i18n.Message{
		ID:    "QueryCardsFilterStatusesPresentation",
		Other: "Status",
	})
}

func CategoriesUrl() string {
	return ApiCardsParams{}.getUrlTagByFieldName("Categories")
}

func CategoriesPresentation(ctx context.Context) string {
	return lang.Message(ctx, &i18n.Message{
		ID:    "QueryCardsFilterCategoriesPresentation",
		Other: "Category",
	})
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
