package model

import "time"

type OptionsGetItems struct {
	PageIndex       int64
	CountForOnePage int64

	SimilarityThreshold float64
	SearchQuery         string

	Categories []Category
	Statuses   []ReleaseStatus
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
	ReleaseStatus        ReleaseStatus
}

type AnimelayerItem struct {
	Id               int64
	Identifier       string
	Title            string
	ReleaseStatus    ReleaseStatus
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
