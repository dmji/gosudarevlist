package model

type FilterItem struct {
	Value string
	Count int64
}

type FilterGroup struct {
	Name  string
	Items []FilterItem
}
