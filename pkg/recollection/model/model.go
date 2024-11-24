package model

import (
	"collector/internal/custom_types"
)

type OptionsGetItems struct {
	PageIndex       int64
	CountForOnePage int64

	SearchQuery string
	Categories  []Category
	IsCompleted custom_types.BoolEx
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
