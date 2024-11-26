package model

type FilterItem struct {
	Presentation string
	Value        string
	Count        int64
}

type FilterGroup struct {
	Name          string
	DisplayTitle  string
	CheckboxItems []FilterItem
}
