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

type UpdateStatus int

const (
	StatusNew UpdateStatus = iota
	StatusRemoved
	StatusUpdated
	StatusUnknown
)

func (s *UpdateStatus) Presentation() string {
	switch *s {
	case StatusNew:
		return "Новое"
	case StatusRemoved:
		return "Удалено"
	case StatusUpdated:
		return "Обновлено"
	default:
		return ""
	}
}
