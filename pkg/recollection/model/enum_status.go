package model

import (
	"context"
	"errors"

	"github.com/dmji/gosudarevlist/lang"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Status string

func StatusFromString(s string) (Status, error) {
	switch s {

	case string(Statuses.Completed):
		return Statuses.Completed, nil
	case string(Statuses.OnAir):
		return Statuses.OnAir, nil
	case string(Statuses.All):
		return Statuses.All, nil
	}

	return Statuses.All, errors.New("string not match any of status")
}

type statuses struct {
	Completed Status
	OnAir     Status
	All       Status
}

// Categories - object to emulate enum class
var Statuses = statuses{
	Completed: "completed",
	OnAir:     "on_air",
}

func (c *Status) Presentation(ctx context.Context) string {
	switch *c {
	case Statuses.Completed:
		return lang.Message(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "ModelReleaseStatusCompletedPresentation",
				Other: "Completed",
			},
		})
	case Statuses.OnAir:
		return lang.Message(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "ModelReleaseStatusOnAirPresentation",
				Other: "On Air",
			},
		})
	case Statuses.All:
		return ""
	default:
		return ""
	}
}
