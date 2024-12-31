package model

//go:generate go-stringer -type=UpdateableField -trimprefix=UpdateableField -output enum_updateable_field_string.go -nametransform=snake_case_lower -fromstringgenfn -outputtransform=snake_case_lower -extraconstsnameprefix=_ -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=pascal_case -extraconstsvaluesuffix=Presentation

import (
	"context"
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

func (c *UpdateableField) Presentation(ctx context.Context) string {
	switch *c {
	case UpdateableFieldTitle:
		return "Title"
	case UpdateableFieldReleaseStatus:
		return "IsCompleted"
	case UpdateableFieldLastCheckedDate:
		return "LastCheckedDate"
	case UpdateableFieldCreatedDate:
		return "CreatedDate"
	case UpdateableFieldUpdatedDate:
		return "UpdatedDate"
	case UpdateableFieldTorrentFilesSize:
		return "TorrentFilesSize"
	case UpdateableFieldNotes:
		return "Notes"
	case UpdateableFieldIdentifier:
		return "Identifier"
	default:
		return ""
	}
}
