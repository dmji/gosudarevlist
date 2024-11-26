package model

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
