package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/internal/updater/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/internal/updater/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"
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

func pgItemFromItem(ctx context.Context, item pgx_sqlc.AnimelayerItem) *model.AnimelayerItem {
	return &model.AnimelayerItem{
		Identifier:       item.Identifier,
		Title:            item.Title,
		ReleaseStatus:    pgxReleaseStatusAnimelayerToReleaseStatusAnimelayer(ctx, item.ReleaseStatus),
		LastCheckedDate:  pgx_utils.TimeFromPgTimestamp(item.LastCheckedDate),
		FirstCheckedDate: pgx_utils.TimeFromPgTimestamp(item.FirstCheckedDate),
		CreatedDate:      pgx_utils.TimeFromPgTimestamp(item.CreatedDate),
		UpdatedDate:      pgx_utils.TimeFromPgTimestamp(item.UpdatedDate),
		RefImageCover:    item.RefImageCover,
		RefImagePreview:  item.RefImagePreview,
		BlobImageCover:   item.BlobImageCover,
		BlobImagePreview: item.BlobImagePreview,
		TorrentFilesSize: item.TorrentFilesSize,
		Notes:            item.Notes,
	}
}
