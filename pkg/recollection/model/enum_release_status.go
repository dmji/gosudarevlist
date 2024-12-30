package model

import (
	"context"
	"errors"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ReleaseStatus string

func (c *ReleaseStatus) String() string {
	return string(*c)
}

var ReleaseStatuses = struct {
	OnAir       ReleaseStatus
	Incompleted ReleaseStatus
	Completed   ReleaseStatus
	All         ReleaseStatus
}{
	OnAir:       "on_air",
	Incompleted: "incompleted",
	Completed:   "completed",
	All:         "",
}

func (c *ReleaseStatus) Presentation(ctx context.Context) string {
	switch *c {
	case ReleaseStatuses.OnAir:
		return lang.Message(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "ModelReleaseStatusOnAirPresentation",
				Other: "On Air",
			},
		})
	case ReleaseStatuses.Incompleted:
		return lang.Message(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "ModelReleaseStatusIncompletedPresentation",
				Other: "Incompleted",
			},
		})
	case ReleaseStatuses.Completed:
		return lang.Message(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "ModelReleaseStatusCompletedPresentation",
				Other: "Completed",
			},
		})
	default:
		return ""
	}
}

func ReleaseStatusFromString(s string) (ReleaseStatus, error) {
	switch s {

	case ReleaseStatuses.OnAir.String():
		return ReleaseStatuses.OnAir, nil
	case ReleaseStatuses.Incompleted.String():
		return ReleaseStatuses.Incompleted, nil
	case ReleaseStatuses.Completed.String():
		return ReleaseStatuses.Completed, nil
	case ReleaseStatuses.All.String():
		return ReleaseStatuses.All, nil
	}

	return ReleaseStatuses.Incompleted, errors.New("string not match any of ReleaseStatusAnimelayer")
}
