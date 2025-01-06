package model

type Page int

type ApiCardsParams struct {
	Page        Page   `qs:"page,omitempty"`
	SearchQuery string `qs:"query,omitempty"`

	Categories []Category      `qs:"category"`
	Statuses   []ReleaseStatus `qs:"release_status"`
}

func WithApiCardsParamsSetPage(p Page) func(*ApiCardsParams) {
	return func(v *ApiCardsParams) {
		v.Page = p
	}
}
