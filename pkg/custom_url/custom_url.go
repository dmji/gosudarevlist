package custom_url

import (
	"net/url"
)

const queryComma = "+"

func QueryCustomParse(query url.Values) url.Values {
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
		qStr += key + "=" + sValues
	}

	q, _ := url.ParseQuery(qStr)

	return q
}

type QueryValuesToStringParam struct {
	bRaw bool
}

type QueryValuesToStringParamOpt func(opt *QueryValuesToStringParam)

func defaultQueryValuesToStringParam() QueryValuesToStringParam {
	return QueryValuesToStringParam{
		bRaw: false,
	}
}

func WithRaw() QueryValuesToStringParamOpt {
	return func(opt *QueryValuesToStringParam) {
		opt.bRaw = true
	}
}

func QueryValuesToString(q *url.Values, opts ...QueryValuesToStringParamOpt) string {
	prm := defaultQueryValuesToStringParam()
	for _, o := range opts {
		o(&prm)
	}

	s := q.Encode()

	if len(s) == 0 {
		return ""
	}
	if prm.bRaw {
		return s
	} else {
		return "?" + s
	}
}
