package model

import (
	"context"
	"fmt"

	"github.com/dmji/gosudarevlist/pkg/enums"
)

type CategoryButton struct {
	category     enums.Category
	urlToCards   string
	urlToUpdates string
	active       bool
}

func (c *CategoryButton) FormatUrlToCards() string {
	return fmt.Sprintf(c.urlToCards, c.category.String())
}

func (c *CategoryButton) FormatUrlToUpdates() string {
	return fmt.Sprintf(c.urlToUpdates, c.category.String())
}

func (c *CategoryButton) Text(ctx context.Context) string {
	return c.category.Presentation(ctx)
}

func (c *CategoryButton) IsActive() bool {
	return c.active
}

func NewCategoryButton(cat enums.Category, urlToCards, urlToUpdates string, active bool) CategoryButton {
	return CategoryButton{
		category:     cat,
		urlToCards:   urlToCards,
		urlToUpdates: urlToUpdates,
		active:       active,
	}
}
