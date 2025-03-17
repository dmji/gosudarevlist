package model

import "github.com/dmji/gosudarevlist/pkg/enums"

type Page int

type ApiCardsParams struct {
	Page        Page   `qs:"page,omitempty"`
	SearchQuery string `qs:"query,omitempty"`

	Statuses []enums.ReleaseStatus `qs:"release_status"`
}

func WithApiCardsParamsSetPage(p Page) func(*ApiCardsParams) {
	return func(v *ApiCardsParams) {
		v.Page = p
	}
}
