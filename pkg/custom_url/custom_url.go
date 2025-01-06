package custom_url

import (
	"net/url"
	"strings"
)

func QueryOrEmpty(s string) string {
	if len(s) == 0 {
		return s
	}
	return "?" + s
}

const queryComma = " "

func QueryCustomParse(s string) (url.Values, error) {

	q, err := url.ParseQuery(s)
	if err != nil {
		return nil, err
	}

	for key := range q {
		vals := make([]string, 0, 5)
		for _, v := range q[key] {
			vals = append(vals, strings.Split(v, queryComma)...)
		}
		q[key] = vals
	}

	return q, nil
}

func QueryCustomEncode(query url.Values) string {
	qStr := ""
	for key, vals := range query {

		sValues := ""
		for _, v := range vals {
			if len(sValues) > 0 && len(v) > 0 {
				sValues += queryComma
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
		qStr += key + "=" + url.QueryEscape(sValues)
	}

	return qStr
}
