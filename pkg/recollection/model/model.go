package model

import "time"

type OptionsGetItems struct {
	Count  int64
	Offset int64

	SearchQuery string
}

type UpdateNote struct {
	ItemID      int64
	UpdateDate  *time.Time
	UpdateTitle string
	ValueOld    string
	ValueNew    string
}
