package handlers

import (
	"io"
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

type WebTheme int8

const (
	WebThemeSystemDefault WebTheme = iota // auto
	WebThemeLight
	WebThemeDark
)

type profileSettings struct {
	language *lang.TagLang
	theme    *WebTheme
}

func (s *router) SettingsHandler(w http.ResponseWriter, r *http.Request) {

	bytedata, _ := io.ReadAll(r.Body)
	reqBodyString := string(bytedata)
	logger.Infow(r.Context(), "SettingsHandler body", "url", reqBodyString)

	settings := profileSettings{}

	const (
		langField = "prefered-language"
	)

	if newValue := settings.language; newValue != nil {
		cookie, err := r.Cookie(langField)
		if err != nil {
			logger.Errorw(r.Context(), "Failed to update cookie value", "name", langField, "error", err)
		} else {
			cookie.Value = newValue.String()
			http.SetCookie(w, cookie)
		}
	}

}
