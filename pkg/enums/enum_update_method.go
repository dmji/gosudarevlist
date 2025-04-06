package enums

//go:generate go run github.com/dmji/go-stringer@latest -type=UpdateMethod -trimprefix=@me -output enum_update_method_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

type UpdateMethod int

const (
	UpdateMethodInsertion UpdateMethod = iota
	UpdateMethodUpdating
)
