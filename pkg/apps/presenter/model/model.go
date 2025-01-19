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
	Description          string
	AnimeLayerRefs       []*ItemCartHrefData
	CategoryPresentation string
	ReleaseStatus        enums.ReleaseStatus
}

type ItemCartHrefData struct {
	Href  string
	Text  []string
	Image string
}
