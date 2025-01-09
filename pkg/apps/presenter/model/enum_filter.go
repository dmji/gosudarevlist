package model

//go:generate go-stringer -type=Filter -trimprefix=Filter -output enum_filter_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

import (
	"context"
	"errors"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Filter int8

const (
	FilterReleaseStatus Filter = iota
	FilterCategory
)

func (c Filter) Presentation(ctx context.Context) string {
	switch c {
	case FilterReleaseStatus:
		return lang.Message(ctx, &i18n.Message{
			ID:    _FilterReleaseStatus_i18n_ID,
			Other: "Release Status",
		})
	case FilterCategory:
		return lang.Message(ctx, &i18n.Message{
			ID:    _FilterCategory_i18n_ID,
			Other: "Category",
		})
	default:
		return ""
	}
}

type presentationStringer interface {
	Presentation(ctx context.Context) string
}

func childPresentationGraber[T presentationStringer](ctx context.Context, s string, FromString func(s string) (T, error)) (string, error) {
	e, err := FromString(s)
	if err != nil {
		return "", err
	}
	return e.Presentation(ctx), nil
}

func (c Filter) ChildsPresentation(ctx context.Context, s string) (string, error) {
	switch c {
	case FilterReleaseStatus:
		return childPresentationGraber(ctx, s, enums.ReleaseStatusFromString)
	case FilterCategory:
		return childPresentationGraber(ctx, s, enums.CategoryFromString)
	default:
		return "", errors.New("wrong category or childs not implemented yet")
	}
}
