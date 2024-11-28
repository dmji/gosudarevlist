package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

type tagLang string

var (
	TagEnglish tagLang = "en"
	TagRussian tagLang = "ru"
)

type Storage struct {
	bundle    *i18n.Bundle
	instances map[tagLang]*Loader
}

func New() *Storage {

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	return &Storage{
		bundle: bundle,
	}
}

func (s *Storage) Reload() {
	s.bundle.MustLoadMessageFile("active.ru.yaml")
}

func (s *Storage) Get(tag tagLang) *Loader {

	res, ok := s.instances[tag]
	if !ok {
		res = &Loader{
			locale: i18n.NewLocalizer(s.bundle, string(tag)),
		}
		s.instances[tag] = res
	}

	return res
}
