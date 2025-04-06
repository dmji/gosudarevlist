package repository_pgx

import (
	"context"

	pgx_sqlc "github.com/dmji/gosudarevlist/internal/presenter/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func releaseStatusAnimelayerToPgxReleaseStatusAnimelayer(ctx context.Context, status enums.ReleaseStatus) pgx_sqlc.ReleaseStatusAnimelayer {
	switch status {

	case enums.ReleaseStatusOnAir:
		return pgx_sqlc.ReleaseStatusAnimelayerOnAir
	case enums.ReleaseStatusIncompleted:
		return pgx_sqlc.ReleaseStatusAnimelayerIncompleted
	case enums.ReleaseStatusCompleted:
		return pgx_sqlc.ReleaseStatusAnimelayerCompleted
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return pgx_sqlc.ReleaseStatusAnimelayerIncompleted
	}
}

func pgxReleaseStatusAnimelayerToReleaseStatusAnimelayer(ctx context.Context, status pgx_sqlc.ReleaseStatusAnimelayer) enums.ReleaseStatus {
	switch status {

	case pgx_sqlc.ReleaseStatusAnimelayerOnAir:
		return enums.ReleaseStatusOnAir
	case pgx_sqlc.ReleaseStatusAnimelayerIncompleted:
		return enums.ReleaseStatusIncompleted
	case pgx_sqlc.ReleaseStatusAnimelayerCompleted:
		return enums.ReleaseStatusCompleted
	default:
		logger.Errorw(ctx, "unexpected model update status", "value", status)
		return enums.ReleaseStatusIncompleted
	}
}

var allReleaseStatus = []pgx_sqlc.ReleaseStatusAnimelayer{
	pgx_sqlc.ReleaseStatusAnimelayerOnAir,
	pgx_sqlc.ReleaseStatusAnimelayerIncompleted,
	pgx_sqlc.ReleaseStatusAnimelayerCompleted,
}

func releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx context.Context, statuses []enums.ReleaseStatus, AppendAllOnEmpty bool) []pgx_sqlc.ReleaseStatusAnimelayer {
	res := make([]pgx_sqlc.ReleaseStatusAnimelayer, 0, len(statuses))

	for _, status := range statuses {
		res = append(res, releaseStatusAnimelayerToPgxReleaseStatusAnimelayer(ctx, status))
	}

	if AppendAllOnEmpty && len(res) == 0 {
		res = append(res, allReleaseStatus...)
	}

	return res
}
