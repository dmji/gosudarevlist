package model

//go:generate go run github.com/dmji/go-stringer@latest -type=CategoryUpdateMode -trimprefix=CategoryUpdateMode -output enum_category_update_mode_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

type CategoryUpdateMode int8

const (
	CategoryUpdateModeWhileNew CategoryUpdateMode = iota
	CategoryUpdateModeAll
)
