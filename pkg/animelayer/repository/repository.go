package animelayer_repository

import (
	sqlc "collector/pkg/animelayer/repository/sqlc"
)

type repository struct {
	query *sqlc.Queries
}

func NewRepository(db sqlc.DBTX) *repository {
	return &repository{
		query: sqlc.New(db),
	}
}
