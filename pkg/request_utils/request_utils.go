package requestutils

import (
	"log"
	"net/http"
)

func LogQuery(r *http.Request, name string) string {
	query := r.URL.Query()
	qStr := ""
	for key, vals := range query {

		sValues := ""
		for i, v := range vals {
			if i != 0 {
				sValues += "-"
			}
			sValues += v
		}

		if len(sValues) == 0 {
			continue
		}

		if len(qStr) == 0 {
			qStr = "?"
		} else {
			qStr += "&"
		}
		qStr += key + "=" + sValues
	}
	log.Printf("Handler | %s: %s", name, qStr)
	return qStr
}
