package filters

type FilterValue struct {
	DisplayName   string
	ParameterName string
	Checked       bool
}

type FilterCategory struct {
	IsCustomThirdStateEnable bool

	DisplayTitle   string
	ParameterTitle string
	Values         []FilterValue
}

type FilterParameters struct {
	Categories  []FilterCategory
	SearchField string
}

func defaultCategories() []FilterCategory {
	return []FilterCategory{
		{
			IsCustomThirdStateEnable: false,
			DisplayTitle:             "Category",
			ParameterTitle:           "category",
			Values: []FilterValue{
				{
					DisplayName:   "Action",
					ParameterName: "action",
					Checked:       false, // slices.ContainsFunc(cats, func(s string) bool { return s == "fantastic" }),
				},
				{
					DisplayName:   "Fantastic",
					ParameterName: "fantastic",
					Checked:       false, // slices.ContainsFunc(cats, func(s string) bool { return s == "fantastic" }),
				},
				{
					DisplayName:   "Historic",
					ParameterName: "historic",
					Checked:       false, // slices.ContainsFunc(cats, func(s string) bool { return s == "historic" }),
				},
				{
					DisplayName:   "Isekai",
					ParameterName: "isekai",
					Checked:       false, // slices.ContainsFunc(cats, func(s string) bool { return s == "isekai" }),
				},
			},
		},
		{
			IsCustomThirdStateEnable: true,
			ParameterTitle:           "st",
			Values: []FilterValue{
				{
					DisplayName:   "On Air",
					ParameterName: "air",
					Checked:       false, // slices.ContainsFunc(sts, func(s string) bool { return s == "air" }),
				},
			},
		},
	}
}

func NewFiltersState(prm *ApiCardsParams) *FilterParameters {
	return &FilterParameters{
		SearchField: prm.SearchQuery,
		/* Categories: []FilterCategory{
			{
				IsCustomThirdStateEnable: false,
				DisplayTitle:             "Category",
				ParameterTitle:           "category",
				Values: []FilterValue{
					{
						DisplayName:   "Action",
						ParameterName: "action",
						Checked:       slices.ContainsFunc(cats, func(s string) bool { return s == "action" }),
					},
					{
						DisplayName:   "Fantastic",
						ParameterName: "fantastic",
						Checked:       slices.ContainsFunc(cats, func(s string) bool { return s == "fantastic" }),
					},
					{
						DisplayName:   "Historic",
						ParameterName: "historic",
						Checked:       slices.ContainsFunc(cats, func(s string) bool { return s == "historic" }),
					},
					{
						DisplayName:   "Isekai",
						ParameterName: "isekai",
						Checked:       slices.ContainsFunc(cats, func(s string) bool { return s == "isekai" }),
					},
				},
			},
			{
				IsCustomThirdStateEnable: true,
				DisplayTitle:             "Status",
				ParameterTitle:           "state",
				Values: []FilterValue{
					{
						DisplayName:   "Completed",
						ParameterName: "completed",
						Checked:       slices.ContainsFunc(sts, func(s string) bool { return s == "completed" }),
					},
				},
			},
		}, */
	}
}
