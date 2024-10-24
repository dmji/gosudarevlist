package filters

import (
	"collector/internal/custom_types"
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

func NewFiltersState(prm *ApiCardsParams) *FilterParameters {
	return &FilterParameters{
		SearchField: prm.SearchQuery,
		Categories: []FilterCategory{
			{
				Parameter: &FilterCategoryParams{
					IsCustomThirdStateEnable: true,
					Name:                     prm.getUrlTagByFieldName("OnAir"),
				},
				Values: []FilterValue{
					{
						DisplayName:   "On Air",
						ParameterName: "air",
						Checked:       prm.OnAir.Value,
					},
				},
			},
		},
	}
}
