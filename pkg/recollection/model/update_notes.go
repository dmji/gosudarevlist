package model

import "time"

type OptionsGetUpdates struct {
	PageIndex       int64
	CountForOnePage int64

	Category Category
}

type UpdateItem struct {
	Date       *time.Time
	Identifier string
	Title      string
	Status     Status
	Notes      []UpdateItemNote
}

type UpdateItemNote struct {
	ValueTitle string
	ValueOld   string
	ValueNew   string
}

type Status int

const (
	StatusNew Status = iota
	StatusRemoved
	StatusUpdated
	StatusUnknown
)
