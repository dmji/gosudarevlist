package model

import "errors"

type Status string

func StatusFromString(s string) (Status, error) {
	switch s {

	case string(Statuses.Completed):
		return Statuses.Completed, nil
	case string(Statuses.OnAir):
		return Statuses.OnAir, nil
	case string(Statuses.All):
		return Statuses.All, nil
	}

	return Statuses.All, errors.New("string not match any of status")
}

type statuses struct {
	Completed Status
	OnAir     Status
	All       Status
}

// Categories - object to emulate enum class
var Statuses = statuses{
	Completed: "completed",
	OnAir:     "on_air",
}

func (c *Status) Presentation() string {
	switch *c {
	case Statuses.Completed:
		return "Completed"
	case Statuses.OnAir:
		return "On Air"
	case Statuses.All:
		return ""
	default:
		return ""
	}
}
