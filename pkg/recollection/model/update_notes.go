package model

import "time"

type OptionsGetUpdates struct {
	PageIndex       int64
	CountForOnePage int64

	Category Category
}

type UpdateNote struct {
	ItemID      int64
	UpdateDate  *time.Time
	UpdateTitle string
	ValueOld    string
	ValueNew    string
}
