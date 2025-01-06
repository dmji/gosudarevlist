package custom_url

import (
	"reflect"
	"strings"

	"github.com/dmji/qs"
)

func init() {
	qs.ApplyOptionsMarshal(qs.WithCustomUrlQueryToStringEncoder(QueryCustomEncode))
	qs.ApplyOptionsUnmarshal(
		qs.WithCustomStringToUrlQueryParser(QueryCustomParse),
		qs.WithCustomSliceToStringFunc(func(a []string) (string, error) {
			return strings.Join(a, queryComma), nil
		}),
	)
	qs.RegisterSubFactoryUnmarshaler(reflect.Slice, newSliceUnmarshaler)
}

func Encode[TIn any](v *TIn) (string, error) {
	s, err := qs.Marshal(v)
	if err != nil {
		return "", err
	}
	return s, nil
}

func Decode[TOut any](s string, inits ...func(*TOut)) (*TOut, error) {
	v := new(TOut)
	for _, init := range inits {
		init(v)
	}

	err := qs.Unmarshal(v, s)
	if err != nil {
		return nil, err
	}
	return v, nil
}
