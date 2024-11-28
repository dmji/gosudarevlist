package lang

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Loader struct {
	locale *i18n.Localizer
}

func (l *Loader) HelloPerson(name string) string {
	return l.locale.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "HelloPerson",
			Description: "Greeting message",
			Other:       "Hello {{.Name}}",
		},
		TemplateData: map[string]string{
			"Name": name,
		},
	})
}
