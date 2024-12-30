package repository_pgx

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func timeToPgTimestamp(t *time.Time) (pgtype.Timestamp, error) {

	pgTime := pgtype.Timestamp{}

	if t != nil {
		if err := pgTime.Scan(*t); err != nil {
			return pgTime, err
		}
	}

	return pgTime, nil
}

func timeFromPgTimestamp(t pgtype.Timestamp) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}
