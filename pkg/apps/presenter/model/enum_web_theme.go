package model

//go:generate go run github.com/dmji/go-stringer@latest -type=WebTheme -trimprefix=WebTheme -output enum_web_theme_string.go -nametransform=snake_case_lower -fromstringgenfn -linecomment -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

type WebTheme int8

const (
	WebThemeSystemDefault WebTheme = iota // auto
	WebThemeLight
	WebThemeDark
)
