package expose_header_utils

import (
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
)

func SetCookiePreferedLanguage(w http.ResponseWriter, value lang.TagLang) *http.Cookie {
	cookie := &http.Cookie{
		Name:   "prefered-language",
		Value:  value.String(),
		Path:   "/",
		MaxAge: 0,
		// HttpOnly: true,
		// Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	return cookie
}

func GetCookiePreferedLanguage(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("prefered-language")
	if err != nil {
		cookie = SetCookiePreferedLanguage(w, lang.TagEnglish)
	}
	return cookie
}
