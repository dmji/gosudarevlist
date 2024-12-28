package model

import "time"

type OptionsGetItems struct {
	PageIndex       int64
	CountForOnePage int64

	SearchQuery string
	Categories  []Category
	Statuses    []Status
}

type ItemCartData struct {
	Title                string
	CreatedDate          string
	UpdatedDate          string
	TorrentWeight        string
	Image                string
	Description          string
	AnimeLayerRef        string
	CategoryPresentation string
	IsCompleted          bool
}

type AnimelayerItem struct {
	Identifier       string
	Title            string
	IsCompleted      bool
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
	Category         Category
}
