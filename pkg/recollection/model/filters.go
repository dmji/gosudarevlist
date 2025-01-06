package model

type FilterItem struct {
	Presentation  string
	Value         string
	Count         int64
	CountFiltered int64
	Selected      bool
}

type FilterGroup struct {
	Name          string
	DisplayTitle  string
	CheckboxItems []FilterItem
}
