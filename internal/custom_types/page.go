package custom_types

import (
	"strconv"
)

type Page int

func (p *Page) UnmarshalText(value []byte) error {
	if len(value) == 0 {
		return nil
	}

	page, err := strconv.ParseInt(string(value), 10, 64)
	if err != nil {
		return err
	}

	*p = Page(page)
	return nil
}
