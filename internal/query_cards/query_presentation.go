package query_cards

import (
	"collector/pkg/logger"
	"collector/pkg/recollection/model"
	"context"
)

type Stringer struct {
	mapItemPresentationFunctions  map[string]func(context.Context, string) string
	mapTitlePresentationFunctions map[string]func(context.Context) string
}

func NewStringer() *Stringer {
	return &Stringer{
		mapItemPresentationFunctions: map[string]func(context.Context, string) string{
			StatusesUrl(): func(ctx context.Context, s string) string {
				v, err := model.StatusFromString(s)
				if err != nil {
					return s
				}
				return v.Presentation(ctx)
			},
			CategoriesUrl(): func(ctx context.Context, s string) string {
				v, err := model.CategoryFromString(s)
				if err != nil {
					return s
				}
				return v.Presentation(ctx)
			},
		},
		mapTitlePresentationFunctions: map[string]func(context.Context) string{
			StatusesUrl():   StatusesPresentation,
			CategoriesUrl(): CategoriesPresentation,
		},
	}
}

func (stringer *Stringer) GetItemPresentation(ctx context.Context, k string, s string) string {
	v, bOk := stringer.mapItemPresentationFunctions[k]
	if !bOk {
		logger.Errorw(ctx, "Not found presentation mapping for key", "key", k)
		return s
	}

	return v(ctx, s)
}

func (stringer *Stringer) GetTitlePresentation(ctx context.Context, k string) string {
	v, bOk := stringer.mapTitlePresentationFunctions[k]
	if !bOk {
		logger.Errorw(ctx, "Not found presentation mapping for title", "title", k)
		return k
	}

	return v(ctx)
}
