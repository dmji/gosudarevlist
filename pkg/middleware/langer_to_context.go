package middleware

import (
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
)

func detectLang(s string) lang.TagLang {
	if s == "ru" {
		return lang.TagRussian
	}

	return lang.TagEnglish
}

func LangerToContextMiddleware(storage *lang.Storage, handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		lang := "en"
		cookie, err := r.Cookie("prefered-language")
		if err == nil {
			lang = cookie.Value
		} else {
			cookie := http.Cookie{
				Name:   "prefered-language",
				Value:  "ru",
				Path:   "/",
				MaxAge: 3600,
				//HttpOnly: true,
				//Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(w, &cookie)
		}

		ctx := storage.ToContext(r.Context(), detectLang(lang))

		handler(w, r.WithContext(ctx))
	}
}
