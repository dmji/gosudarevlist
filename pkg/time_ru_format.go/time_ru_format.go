package time_ru_format

import (
	"fmt"
	"time"
)

func monthToRussian(month time.Month) string {
	switch month {
	case time.January:
		return "Января"
	case time.February:
		return "Февраля"
	case time.March:
		return "Марта"
	case time.April:
		return "Апреля"
	case time.May:
		return "Мая"
	case time.June:
		return "Июня"
	case time.July:
		return "Июля"
	case time.August:
		return "Августа"
	case time.September:
		return "Сентября"
	case time.October:
		return "Октября"
	case time.November:
		return "Ноября"
	case time.December:
		return "Декабря"
	}
	return ""
}

func Format(t *time.Time) string {
	if t == nil {
		return ""
	}

	if t.Year() == time.Now().Year() {
		return fmt.Sprintf(t.Format("02 %s %s 15:04"), monthToRussian(t.Month()), "в")
	}

	return fmt.Sprintf(t.Format("02 %s 2006 %s 15:04"), monthToRussian(t.Month()), "в")
}
