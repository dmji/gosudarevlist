package custom_types

import (
	"strconv"
)

type Page int

func (p *Page) DecodeValues(value string, defaultValue int) error {
	*p = Page(defaultValue)
	if len(value) == 0 {
		return nil
	}

	page, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}

	*p = Page(page)
	return nil
}
