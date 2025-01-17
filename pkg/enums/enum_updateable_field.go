package enums

//go:generate go run github.com/dmji/go-stringer@latest -type=UpdateableField -trimprefix=UpdateableField -output enum_updateable_field_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

import (
	"context"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UpdateableField int8

const (
	UpdateableFieldTitle UpdateableField = iota
	UpdateableFieldReleaseStatus
	UpdateableFieldLastCheckedDate
	UpdateableFieldCreatedDate
	UpdateableFieldUpdatedDate
	UpdateableFieldTorrentFilesSize
	UpdateableFieldNotes
	UpdateableFieldIdentifier
)

func (c UpdateableField) Presentation(ctx context.Context) string {
	switch c {

	case UpdateableFieldTitle:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateableFieldTitle_i18n_ID,
			Other: "Title",
		})
	case UpdateableFieldReleaseStatus:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateableFieldReleaseStatus_i18n_ID,
			Other: "Release Status",
		})
	case UpdateableFieldLastCheckedDate:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateableFieldLastCheckedDate_i18n_ID,
			Other: "LastCheckedDate",
		})
	case UpdateableFieldCreatedDate:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateableFieldCreatedDate_i18n_ID,
			Other: "CreatedDate",
		})
	case UpdateableFieldUpdatedDate:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateableFieldUpdatedDate_i18n_ID,
			Other: "UpdatedDate",
		})
	case UpdateableFieldTorrentFilesSize:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateableFieldTorrentFilesSize_i18n_ID,
			Other: "TorrentFilesSize",
		})
	case UpdateableFieldNotes:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateableFieldNotes_i18n_ID,
			Other: "Notes",
		})
	case UpdateableFieldIdentifier:
		return lang.Message(ctx, &i18n.Message{
			ID:    _UpdateableFieldIdentifier_i18n_ID,
			Other: "Identifier",
		})
	default:
		return ""
	}
}
