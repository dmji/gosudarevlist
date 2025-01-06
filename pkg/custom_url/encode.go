package custom_url

import (
	"fmt"
	"reflect"

	"github.com/dmji/qs"
)

func init() {
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

type sliceUnmarshaler struct {
	Type            reflect.Type
	ElemUnmarshaler qs.Unmarshaler
}

func newSliceUnmarshaler(t reflect.Type, opts *qs.UnmarshalOptions) (qs.Unmarshaler, error) {
	if t.Kind() != reflect.Slice {
		return nil, &wrongKindError{Expected: reflect.Slice, Actual: t}
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
		return &wrongTypeError{Actual: t, Expected: p.Type}
	}

	if v.IsNil() {
		v.Set(reflect.MakeSlice(t, len(a), len(a)))
	}

	n := 0
	for i := range a {
		err := p.ElemUnmarshaler.Unmarshal(v.Index(i), a[i:i+1], opts)
		if err == nil {
			n++
		}
	}
	v.Set(v.Slice(0, n))

	return nil
}

type wrongTypeError struct {
	Actual   reflect.Type
	Expected reflect.Type
}

func (e *wrongTypeError) Error() string {
	return fmt.Sprintf("received type %v, want %v", e.Actual, e.Expected)
}

type wrongKindError struct {
	Actual   reflect.Type
	Expected reflect.Kind
}

func (e *wrongKindError) Error() string {
	return fmt.Sprintf("received type %v of kind %v, want kind %v",
		e.Actual, e.Actual.Kind(), e.Expected)
}
