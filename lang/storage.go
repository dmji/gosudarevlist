package lang

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

type TagLang string

func (c *TagLang) String() string {
	return string(*c)
}

var (
	TagEnglish TagLang = "en"
	TagRussian TagLang = "ru"
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
