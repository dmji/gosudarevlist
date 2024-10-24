package custom_types

import (
	"net/url"
)

type BoolEx int

type BoolExProp struct {
	Value BoolEx
	Name  string
}

func (b *BoolEx) Bool() bool {
	if *b == BoolExFalse {
		return false
	}
	return true
}

// EncodeValues - encoder
func (b *BoolExProp) EncodeValues(key string, v *url.Values) error {
	switch b.Value {
	case BoolExTrue:
		v.Add(key, b.Name)
		return nil
	case BoolExIntermediate:
		v.Add(key, "~"+b.Name)
		return nil
	default:
		return nil
	}
}

func (b *BoolExProp) DecodeValues(value string) error {
	switch value {
	case b.Name:
		b.Value = BoolExTrue
		return nil
	case "~" + b.Name:
		b.Value = BoolExIntermediate
		return nil
	default:
		b.Value = BoolExFalse
		return nil
	}
}

const (
	BoolExFalse BoolEx = iota
	BoolExIntermediate
	BoolExTrue
)
