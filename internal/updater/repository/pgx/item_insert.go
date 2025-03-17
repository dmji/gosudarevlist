package repository_pgx

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dmji/gosudarevlist/internal/updater/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/internal/updater/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"

	"github.com/jackc/pgx/v5"
)

func (repo *repository) InsertItem(ctx context.Context, item *model.AnimelayerItem, category enums.Category) error {
	now := time.Now()
	lastCheckedDate, err := pgx_utils.TimeToPgTimestamp(&now)
	if err != nil {
		return err
	}

	createdDate, err := pgx_utils.TimeToPgTimestamp(item.CreatedDate)
	if err != nil {
		return err
	}

	updatedDate, err := pgx_utils.TimeToPgTimestamp(item.UpdatedDate)
	if err != nil {
		return err
	}

	//
	// Start transaction
	//
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	r := repo.query.WithTx(tx)

	itemId, err := r.InsertItem(ctx,
		pgx_sqlc.InsertItemParams{
			Identifier:       item.Identifier,
			Title:            item.Title,
			ReleaseStatus:    releaseStatusAnimelayerToPgxReleaseStatusAnimelayer(ctx, item.ReleaseStatus),
			LastCheckedDate:  lastCheckedDate,
			CreatedDate:      createdDate,
			UpdatedDate:      updatedDate,
			RefImageCover:    item.RefImageCover,
			RefImagePreview:  item.RefImagePreview,
			BlobImageCover:   "",
			BlobImagePreview: "",
			TorrentFilesSize: item.TorrentFilesSize,
			Notes:            item.Notes,
			Category:         categoriesToAnimelayerCategory(category),
		},
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}

	err = repo.InsertUpdateNote(ctx, model.UpdateItem{
		Date:         &now,
		UpdateStatus: enums.UpdateStatusNew,
		Notes:        []model.UpdateItemNote{},
		ItemId:       itemId,
		// Identifier:   item.Identifier,
	})
	if err != nil {
		return err
	}

	//
	// Commit
	//
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
