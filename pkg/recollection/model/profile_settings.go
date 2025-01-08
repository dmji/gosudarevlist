package model

import "github.com/dmji/gosudarevlist/lang"

type ProfileSettings struct {
	Language *lang.TagLang `qs:"language,omitempty,nil"`
	Theme    *WebTheme     `qs:"theme,omitempty,nil"`
}
