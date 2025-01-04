package query_cards

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

type Stringer struct {
	mapItemPresentationFunctions  map[string]func(context.Context, string) string
	mapTitlePresentationFunctions map[string]func(context.Context) string
}

type presentationStringer interface {
	Presentation(ctx context.Context) string
}

func presentationFromString[T presentationStringer](enumFromString func(string) (T, error)) func(context.Context, string) string {
	return func(ctx context.Context, s string) string {
		v, err := enumFromString(s)
		if err != nil {
			return s
		}
		return v.Presentation(ctx)
	}
}

func NewStringer() *Stringer {
	return &Stringer{
		mapItemPresentationFunctions: map[string]func(context.Context, string) string{
			StatusesUrl():   presentationFromString(model.ReleaseStatusFromString),
			CategoriesUrl(): presentationFromString(model.CategoryFromString),
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
