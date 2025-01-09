package model

import (
	"time"

	"github.com/dmji/gosudarevlist/pkg/enums"
)

type AnimelayerItem struct {
	Id               int64
	Identifier       string
	Title            string
	ReleaseStatus    enums.ReleaseStatus
	LastCheckedDate  *time.Time
	FirstCheckedDate *time.Time
	CreatedDate      *time.Time
	UpdatedDate      *time.Time
	RefImageCover    string
	RefImagePreview  string
	BlobImageCover   string
	BlobImagePreview string
	TorrentFilesSize string
	Notes            string
	Category         enums.Category
}
