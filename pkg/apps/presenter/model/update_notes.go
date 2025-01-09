package model

import (
	"time"

	"github.com/dmji/gosudarevlist/pkg/enums"
)

type UpdateItemNote struct {
	ValueTitle enums.UpdateableField
	ValueOld   string
	ValueNew   string
}

type UpdateItem struct {
	Date         *time.Time
	Title        string
	UpdateStatus enums.UpdateStatus
	Notes        []UpdateItemNote
}
