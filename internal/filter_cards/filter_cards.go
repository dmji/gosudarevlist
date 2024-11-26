package filter_cards

/*
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
				DisplayTitle: "Status",
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
			{
				DisplayTitle: "Categories",
				Parameter: &FilterCategoryParams{
					IsCustomThirdStateEnable: false,
					Name:                     prm.CategoriesUrl(),
				},
				Values: []FilterValue{
					{
						DisplayName:   "Completed",
						ParameterName: prm.IsCompleted.Name,
						Checked:       prm.IsCompleted.Value,
					},
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
*/
