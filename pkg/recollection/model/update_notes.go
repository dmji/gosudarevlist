package model

import "time"

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
