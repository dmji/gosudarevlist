package model

//go:generate go-stringer -type=ReleaseStatus -trimprefix=ReleaseStatus -output enum_release_status_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

import (
	"context"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ReleaseStatus int8

const (
	ReleaseStatusOnAir ReleaseStatus = iota
	ReleaseStatusIncompleted
	ReleaseStatusCompleted
)

func (c ReleaseStatus) Presentation(ctx context.Context) string {
	switch c {
	case ReleaseStatusOnAir:
		return lang.Message(ctx, &i18n.Message{
			ID:    _ReleaseStatusOnAir_i18n_ID,
			Other: "On Air",
		})
	case ReleaseStatusIncompleted:
		return lang.Message(ctx, &i18n.Message{
			ID:    _ReleaseStatusIncompleted_i18n_ID,
			Other: "Incompleted",
		})
	case ReleaseStatusCompleted:
		return lang.Message(ctx, &i18n.Message{
			ID:    _ReleaseStatusCompleted_i18n_ID,
			Other: "Completed",
		})
	default:
		return ""
	}
}
