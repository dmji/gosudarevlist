package model

import (
	"context"
	"errors"
)

//go:generate go run t~/t.go -type=UpdateableField

type UpdateableField string

func (c *UpdateableField) String() string {
	return string(*c)
}

var UpdateableFields = struct {
	Title            UpdateableField
	ReleaseStatus    UpdateableField
	LastCheckedDate  UpdateableField
	CreatedDate      UpdateableField
	UpdatedDate      UpdateableField
	TorrentFilesSize UpdateableField
	Notes            UpdateableField
	Identifier       UpdateableField
}{
	Title:            "title",
	ReleaseStatus:    "release_status",
	LastCheckedDate:  "last_checked_date",
	CreatedDate:      "created_date",
	UpdatedDate:      "updated_date",
	TorrentFilesSize: "torrent_files_size",
	Notes:            "notes",
	Identifier:       "identifier",
}

func (c *UpdateableField) Presentation(ctx context.Context) string {
	switch *c {
	case UpdateableFields.Title:
		return "Title"
	case UpdateableFields.ReleaseStatus:
		return "IsCompleted"
	case UpdateableFields.LastCheckedDate:
		return "LastCheckedDate"
	case UpdateableFields.CreatedDate:
		return "CreatedDate"
	case UpdateableFields.UpdatedDate:
		return "UpdatedDate"
	case UpdateableFields.TorrentFilesSize:
		return "TorrentFilesSize"
	case UpdateableFields.Notes:
		return "Notes"
	case UpdateableFields.Identifier:
		return "Identifier"
	default:
		return ""
	}
}

func UpdateableFieldFromString(s string) (UpdateableField, error) {
	switch s {

	case UpdateableFields.Title.String():
		return UpdateableFields.Title, nil
	case UpdateableFields.ReleaseStatus.String():
		return UpdateableFields.ReleaseStatus, nil
	case UpdateableFields.LastCheckedDate.String():
		return UpdateableFields.LastCheckedDate, nil
	case UpdateableFields.CreatedDate.String():
		return UpdateableFields.CreatedDate, nil
	case UpdateableFields.UpdatedDate.String():
		return UpdateableFields.UpdatedDate, nil
	case UpdateableFields.TorrentFilesSize.String():
		return UpdateableFields.TorrentFilesSize, nil
	case UpdateableFields.Notes.String():
		return UpdateableFields.Notes, nil
	case UpdateableFields.Identifier.String():
		return UpdateableFields.Identifier, nil
	}

	return UpdateableFields.Title, errors.New("string not match any of UpdateableFields")
}
