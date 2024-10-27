package filter_cards

import (
	"collector/internal/custom_types"
	"collector/internal/query_cards"
)

type FilterValue struct {
	DisplayName   string
	ParameterName string
	Checked       custom_types.BoolEx
}

type FilterCategoryParams struct {
	IsCustomThirdStateEnable bool

	Name string
}

type FilterCategory struct {
	DisplayTitle string
	Parameter    *FilterCategoryParams
	Values       []FilterValue
}

type FilterParameters struct {
	Categories  []FilterCategory
	SearchField string
}

func NewFiltersState(prm *query_cards.ApiCardsParams) *FilterParameters {
	return &FilterParameters{
		SearchField: prm.SearchQuery,
		Categories: []FilterCategory{
			{
				Parameter: &FilterCategoryParams{
					IsCustomThirdStateEnable: true,
					Name:                     prm.IsCompletedUrl(),
				},
				Values: []FilterValue{
					{
						DisplayName:   "Completed",
						ParameterName: prm.IsCompleted.Name,
						Checked:       prm.IsCompleted.Value,
					},
				},
			},
		},
	}
}