package service

import (
	"fmt"
	"time"

	"github.com/dmji/gosudarevlist/pkg/enums"
)

type errorInProcess struct {
	category enums.Category
	time     *time.Time
}

func (e *errorInProcess) Error() string {
	return fmt.Sprintf("updater for category '%s' already in progress from %s", e.category, e.time.UTC().String())
}

func NewRrrorInProcess(cat enums.Category, t *time.Time) error {
	return &errorInProcess{
		time:     t,
		category: cat,
	}
}

func IsErrorInProcess(e error) (*time.Time, bool) {
	if re, ok := e.(*errorInProcess); ok {
		return re.time, true
	}
	return nil, false
}
