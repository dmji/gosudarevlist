package model

import (
	"collector/internal/custom_types"
	"time"
)

type OptionsGetNotes struct {
	Count  int64
	Offset int64
}

type OptionsGetItems struct {
	Count  int64
	Offset int64

	SearchQuery string
	Category    Category
	IsCompleted custom_types.BoolEx
}

type UpdateNote struct {
	ItemID      int64
	UpdateDate  *time.Time
	UpdateTitle string
	ValueOld    string
	ValueNew    string
}

type ItemCartData struct {
	Title         string
	CreatedDate   string
	UpdatedDate   string
	TorrentWeight string
	Image         string
	Description   string
	AnimeLayerRef string
	IsCompleted   bool
}
