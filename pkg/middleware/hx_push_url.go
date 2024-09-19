package middleware

import (
	"log"
	"net/http"
	"net/url"
)

func queryCustomParse(r *http.Request) url.Values {
	query := r.URL.Query()
	qStr := ""
	for key, vals := range query {

		sValues := ""
		for i, v := range vals {
			if i != 0 && len(v) > 0 {
				sValues += "-"
			}
			sValues += v
		}

		if len(sValues) == 0 {
			continue
		}

		if len(qStr) == 0 {
		} else {
			qStr += "&"
		}
		qStr += key + "=" + sValues
	}

	q, _ := url.ParseQuery(qStr)

	return q
}

func queryValuesToString(q *url.Values) string {

	q.Del("page")

	s := q.Encode()

	if len(s) == 0 {
		return ""
	}

	return "?" + s
}

func HxPushUrlMiddleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		queryValues := queryCustomParse(r)
		//log.Printf("Handler | %s: %s", name, qStr)

		currentUri := r.Header.Get("HX-Current-URL")
		currentUrl, err := url.Parse(currentUri)
		if err != nil {
			log.Panic(err)
		}

		newUri := currentUrl.Path + queryValuesToString(&queryValues)

		w.Header().Set("Access-Control-Expose-Headers", "Hx-Push-Url")
		w.Header().Set("Hx-Push-Url", newUri)

		log.Printf("Middleware Hx-Push-Url | New URI: %s", newUri)

		handler(w, r)
	}
}
