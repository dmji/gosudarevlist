package time_formater

import (
	"context"
	"fmt"
	"time"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func monthToRussian(ctx context.Context, month time.Month) string {
	switch month {
	case time.January:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthJanuary", Other: "January"})
	case time.February:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthFebruary", Other: "February"})
	case time.March:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthMarch", Other: "March"})
	case time.April:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthApril", Other: "April"})
	case time.May:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthMay", Other: "May"})
	case time.June:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthJune", Other: "June"})
	case time.July:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthJuly", Other: "July"})
	case time.August:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthAugust", Other: "August"})
	case time.September:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthSeptember", Other: "September"})
	case time.October:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthOctober", Other: "October"})
	case time.November:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthNovember", Other: "November"})
	case time.December:
		return lang.Message(ctx, &i18n.Message{ID: "TimeFormatMonthDecember", Other: "December"})
	}
	return ""
}

func inWord(ctx context.Context) string {
	return lang.Message(ctx, &i18n.Message{
		ID:    "TimeFormatInWord",
		Other: "in",
	})
}

func Format(ctx context.Context, t *time.Time) string {
	if t == nil {
		return ""
	}

	if t.Year() == time.Now().Year() {
		return fmt.Sprintf(t.Format("02 %s %s 15:04"), monthToRussian(ctx, t.Month()), inWord(ctx))
	}

	return fmt.Sprintf(t.Format("02 %s 2006 %s 15:04"), monthToRussian(ctx, t.Month()), inWord(ctx))
}
