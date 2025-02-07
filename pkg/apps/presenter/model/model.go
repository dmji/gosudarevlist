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
	AnimeLayerRefs       []*ItemCartHrefData
	CategoryPresentation string
	ReleaseStatus        enums.ReleaseStatus
}

type ItemCartDescriptions struct {
	Type            string   //"Тип:"
	Genres          []string // "Жанр"
	Year            string   //"Год выхода"
	EpisodeCount    string   //"Кол серий"
	EpisodeDuration string   //"Продолжительность"
	UpdateReaseon   string   //"Торрент был обновлен"
}

type ItemCartHrefData struct {
	Href        string
	Text        []string
	Image       string
	Description ItemCartDescriptions
}
