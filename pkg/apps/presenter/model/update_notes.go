package model

import "time"

type UpdateItemNote struct {
	ValueTitle UpdateableField
	ValueOld   string
	ValueNew   string
}

type UpdateItem struct {
	Date         *time.Time
	Title        string
	UpdateStatus UpdateStatus
	Notes        []UpdateItemNote

	ItemId     int64
	Identifier string
}
