package repository_pgx

import (
	"context"
	"time"

	"github.com/dmji/gosudarevlist/internal/updater/model"
	r "github.com/dmji/gosudarevlist/internal/updater/repository"
	pgx_sqlc "github.com/dmji/gosudarevlist/internal/updater/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/time_formater.go"

	"github.com/jackc/pgx/v5"
)

func (repo *repository) UpdateItem(ctx context.Context, item *model.AnimelayerItem) error {
	now := time.Now()
	oldItem, err := repo.GetItemByIdentifier(ctx, item.Identifier)
	if err != nil {
		return err
	}

	//
	// Collect updated notes
	//
	arg, notes := compareItems(ctx, oldItem, item)
	if arg == nil {
		return r.NewErrorItemNotChanged(item.Identifier)
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

	itemId, err := r.UpdateItem(ctx, *arg)
	if err != nil {
		return err
	}

	//
	// Push updated notes
	//
	if len(notes) > 0 {
		err = repo.InsertUpdateNote(ctx, model.UpdateItem{
			Date:         &now,
			UpdateStatus: enums.UpdateStatusUpdated,
			Notes:        notes,
			ItemId:       itemId,
			// Identifier:   item.Identifier,
		})
		if err != nil {
			return err
		}
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

func compareItems(ctx context.Context, oldItem, item *model.AnimelayerItem) (*pgx_sqlc.UpdateItemParams, []model.UpdateItemNote) {
	bUpdateRequied := false
	itemUpdate := &pgx_sqlc.UpdateItemParams{
		Identifier: oldItem.Identifier,
	}
	itemUpdate.LastCheckedDate.Scan(*item.LastCheckedDate)

	itemNotes := make([]model.UpdateItemNote, 0, 10)

	if isDiffString(oldItem.Title, item.Title) {
		itemUpdate.Title.Scan(item.Title)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: enums.UpdateableFieldTitle,
			ValueOld:   oldItem.Title,
			ValueNew:   item.Title,
		})
		bUpdateRequied = true
	}

	if isDiffString(oldItem.ReleaseStatus.String(), item.ReleaseStatus.String()) {
		itemUpdate.ReleaseStatus.Scan(item.ReleaseStatus.String())
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: enums.UpdateableFieldReleaseStatus,
			ValueOld:   oldItem.ReleaseStatus.Presentation(ctx),
			ValueNew:   item.ReleaseStatus.Presentation(ctx),
		})
		bUpdateRequied = true
	}

	if isDiffTimes(oldItem.CreatedDate, item.CreatedDate) {
		itemUpdate.CreatedDate.Scan(*item.CreatedDate)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: enums.UpdateableFieldCreatedDate,
			ValueOld:   time_formater.Format(ctx, oldItem.CreatedDate),
			ValueNew:   time_formater.Format(ctx, item.CreatedDate),
		})
		bUpdateRequied = true
	}

	if isDiffTimes(oldItem.UpdatedDate, item.UpdatedDate) {
		itemUpdate.UpdatedDate.Scan(*item.UpdatedDate)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: enums.UpdateableFieldUpdatedDate,
			ValueOld:   time_formater.Format(ctx, oldItem.UpdatedDate),
			ValueNew:   time_formater.Format(ctx, item.UpdatedDate),
		})
		bUpdateRequied = true
	}

	if isDiffString(oldItem.RefImageCover, item.RefImageCover) {
		itemUpdate.RefImageCover.Scan(item.RefImageCover)
		bUpdateRequied = true
	}

	if isDiffString(oldItem.RefImagePreview, item.RefImagePreview) {
		itemUpdate.RefImagePreview.Scan(item.RefImagePreview)
		bUpdateRequied = true
	}

	if isDiffString(oldItem.BlobImageCover, item.BlobImageCover) {
		itemUpdate.BlobImageCover.Scan(item.BlobImageCover)
		bUpdateRequied = true
	}

	if isDiffString(oldItem.BlobImagePreview, item.BlobImagePreview) {
		itemUpdate.BlobImagePreview.Scan(item.BlobImagePreview)
		bUpdateRequied = true
	}

	if isDiffString(oldItem.TorrentFilesSize, item.TorrentFilesSize) {
		itemUpdate.TorrentFilesSize.Scan(item.TorrentFilesSize)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: enums.UpdateableFieldTorrentFilesSize,
			ValueOld:   oldItem.TorrentFilesSize,
			ValueNew:   item.TorrentFilesSize,
		})
		bUpdateRequied = true
	}

	if isDiffString(oldItem.Notes, item.Notes) {
		itemUpdate.Notes.Scan(item.Notes)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: enums.UpdateableFieldNotes,
			ValueOld:   oldItem.Notes,
			ValueNew:   item.Notes,
		})
		bUpdateRequied = true
	}

	if bUpdateRequied == false {
		return nil, nil
	}
	return itemUpdate, itemNotes
}

func isDiffString(oldItem, item string) bool {
	if len(oldItem) == 0 && len(item) > 0 {
		return true
	}

	return oldItem != item
}

func isDiffTimes(oldItem, item *time.Time) bool {
	if item == nil {
		return false
	}

	if oldItem == nil {
		return true
	}

	return *oldItem != *item
}
