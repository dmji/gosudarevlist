package model

//go:generate go-stringer -type=UpdateStatus -trimprefix=UpdateStatus -output enum_update_status_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

import (
	"context"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UpdateStatus int

const (
	UpdateStatusNew UpdateStatus = iota
	UpdateStatusRemoved
	UpdateStatusUpdated
	UpdateStatusUnknown
)

func (c UpdateStatus) Presentation(ctx context.Context) string {
	switch c {
	case UpdateStatusNew:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateStatusNew_i18n_ID,
			Other: "On Air",
		})
	case UpdateStatusRemoved:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateStatusRemoved_i18n_ID,
			Other: "Removed",
		})
	case UpdateStatusUpdated:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateStatusUpdated_i18n_ID,
			Other: "Updated",
		})
	default:
		return ""
	}
}
