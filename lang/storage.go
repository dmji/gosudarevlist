package lang

//go:generate go-stringer -type=TagLang -trimprefix=TagLang -output enum_tag_lang_string.go -nametransform=snake_case_lower -fromstringgenfn -linecomment -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

import (
	"context"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

type langerCtx string

const (
	langerCtxValue langerCtx = "langer"
)

type TagLang int8

const (
	TagEnglish TagLang = iota // en
	TagRussian                // ru
)

type Storage struct {
	bundle    *i18n.Bundle
	instances map[TagLang]*Loader
}

func New() *Storage {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	s := &Storage{
		bundle:    bundle,
		instances: make(map[TagLang]*Loader),
	}

	s.Reload()

	return s
}

func (s *Storage) Get(tag TagLang) *Loader {
	res, ok := s.instances[tag]
	if !ok {
		res = &Loader{
			locale: i18n.NewLocalizer(s.bundle, tag.String()),
			Tag:    tag,
		}
		s.instances[tag] = res
	}

	return res
}

func (s *Storage) ToContext(ctx context.Context, tag TagLang) context.Context {
	return context.WithValue(ctx, langerCtxValue, s.Get(tag))
}

func FromContext(ctx context.Context) *Loader {
	if loggerC, ok := ctx.Value(langerCtxValue).(*Loader); ok {
		return loggerC
	}

	return nil
}
