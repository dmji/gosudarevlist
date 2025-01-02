package lang

import (
	"context"
	"fmt"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Loader struct {
	locale *i18n.Localizer
	Tag    TagLang
}

func Message(ctx context.Context, cfg *i18n.Message) string {
	loader := FromContext(ctx)
	if loader == nil {
		logger.Errorw(ctx, "Localize storage provider error", "error", "string loader not found")
		return "string loader not found"
	}

	str, err := loader.locale.LocalizeMessage(cfg)
	if err != nil {
		logger.Infow(ctx, "Localize error", "warning", fmt.Errorf("not found localization to '%s' for ID '%s'", loader.Tag, cfg.ID))
	}
	return str
}

func MustLocalize(ctx context.Context, cfg *i18n.LocalizeConfig) string {
	loader := FromContext(ctx)
	if loader == nil {
		logger.Errorw(ctx, "Localize storage provider error", "error", "string loader not found")
		return "string loader not found"
	}

	return loader.locale.MustLocalize(cfg)
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
