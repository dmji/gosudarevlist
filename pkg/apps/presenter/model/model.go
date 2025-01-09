package model

import "github.com/dmji/gosudarevlist/pkg/enums"

type OptionsGetItems struct {
	PageIndex       int64
	CountForOnePage int64

	SimilarityThreshold float64
	SearchQuery         string

	Categories []enums.Category
	Statuses   []enums.ReleaseStatus
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
	ReleaseStatus        enums.ReleaseStatus
}
