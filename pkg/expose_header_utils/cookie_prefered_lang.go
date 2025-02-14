package expose_header_utils

import (
	"context"
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func SetCookiePreferedLanguage(ctx context.Context, w http.ResponseWriter, value lang.TagLang) *http.Cookie {
	if false {
		logger.Errorw(ctx, "SetCookiePreferedLanguage | Cookie updating", "value", value.String())
	}
	cookie := &http.Cookie{
		Name:   "prefered-language",
		Value:  value.String(),
		Path:   "/",
		MaxAge: 3600,
		// HttpOnly: true,
		// Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
	return cookie
}

func GetCookiePreferedLanguage(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("prefered-language")
	if err != nil {
		if false {
			logger.Errorw(r.Context(), "GetCookiePreferedLanguage | Cookie not found, create new value")
		}
		cookie = SetCookiePreferedLanguage(r.Context(), w, lang.TagEnglish)
	}
	return cookie
}
