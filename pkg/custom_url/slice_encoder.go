package custom_url

import (
	"reflect"

	"github.com/dmji/qs"
)

type sliceUnmarshaler struct {
	Type            reflect.Type
	ElemUnmarshaler qs.Unmarshaler
}

func newSliceUnmarshaler(t reflect.Type, opts *qs.UnmarshalOptions) (qs.Unmarshaler, error) {
	if t.Kind() != reflect.Slice {
		return nil, &qs.WrongKindError{Expected: reflect.Slice, Actual: t}
	}

	eu, err := opts.UnmarshalerFactory.Unmarshaler(t.Elem(), opts)
	if err != nil {
		return nil, err
	}

	return &sliceUnmarshaler{
		Type:            t,
		ElemUnmarshaler: eu,
	}, nil
}

func (p *sliceUnmarshaler) Unmarshal(v reflect.Value, a []string, opts *qs.UnmarshalOptions) error {
	t := v.Type()
	if t != p.Type {
		return &qs.WrongTypeError{Actual: t, Expected: p.Type}
	}

	if v.IsNil() {
		v.Set(reflect.MakeSlice(t, len(a), len(a)))
	}

	n := 0
	for i := range a {
		err := p.ElemUnmarshaler.Unmarshal(v.Index(n), a[i:i+1], opts)
		if err == nil {
			n++
		}
	}
	v.Set(v.Slice(0, n))

	return nil
}
