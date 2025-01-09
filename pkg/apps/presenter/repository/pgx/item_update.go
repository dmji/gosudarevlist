package repository_pgx

import (
	"context"
	"time"

	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/time_ru_format.go"

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

	err = repo.InsertUpdateNote(ctx, model.UpdateItem{
		Date:         &now,
		UpdateStatus: model.UpdateStatusUpdated,
		Notes:        notes,
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

func compareItems(ctx context.Context, oldItem, item *model.AnimelayerItem) (*pgx_sqlc.UpdateItemParams, []model.UpdateItemNote) {
	itemUpdate := &pgx_sqlc.UpdateItemParams{
		Identifier: oldItem.Identifier,
	}
	itemNotes := make([]model.UpdateItemNote, 0, 10)

	if isDiffString(oldItem.Title, item.Title) {
		itemUpdate.Title.Scan(item.Title)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: model.UpdateableFieldTitle,
			ValueOld:   oldItem.Title,
			ValueNew:   item.Title,
		})
	}

	if isDiffString(oldItem.ReleaseStatus.String(), item.ReleaseStatus.String()) {
		itemUpdate.ReleaseStatus.Scan(item.ReleaseStatus.String())
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: model.UpdateableFieldReleaseStatus,
			ValueOld:   oldItem.ReleaseStatus.Presentation(ctx),
			ValueNew:   item.ReleaseStatus.Presentation(ctx),
		})
	}

	if isDiffTimes(oldItem.CreatedDate, item.CreatedDate) {
		itemUpdate.CreatedDate.Scan(*item.CreatedDate)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: model.UpdateableFieldCreatedDate,
			ValueOld:   time_ru_format.Format(oldItem.CreatedDate),
			ValueNew:   time_ru_format.Format(item.CreatedDate),
		})
	}

	if isDiffTimes(oldItem.UpdatedDate, item.UpdatedDate) {
		itemUpdate.UpdatedDate.Scan(*item.UpdatedDate)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: model.UpdateableFieldUpdatedDate,
			ValueOld:   time_ru_format.Format(oldItem.UpdatedDate),
			ValueNew:   time_ru_format.Format(item.UpdatedDate),
		})
	}

	if isDiffTimes(oldItem.LastCheckedDate, item.LastCheckedDate) {
		itemUpdate.LastCheckedDate.Scan(*item.LastCheckedDate)
	}

	if isDiffString(oldItem.RefImageCover, item.RefImageCover) {
		itemUpdate.RefImageCover.Scan(item.RefImageCover)
	}

	if isDiffString(oldItem.RefImagePreview, item.RefImagePreview) {
		itemUpdate.RefImagePreview.Scan(item.RefImagePreview)
	}

	if isDiffString(oldItem.BlobImageCover, item.BlobImageCover) {
		itemUpdate.BlobImageCover.Scan(item.BlobImageCover)
	}

	if isDiffString(oldItem.BlobImagePreview, item.BlobImagePreview) {
		itemUpdate.BlobImagePreview.Scan(item.BlobImagePreview)
	}

	if isDiffString(oldItem.TorrentFilesSize, item.TorrentFilesSize) {
		itemUpdate.TorrentFilesSize.Scan(item.TorrentFilesSize)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: model.UpdateableFieldTorrentFilesSize,
			ValueOld:   oldItem.TorrentFilesSize,
			ValueNew:   item.TorrentFilesSize,
		})
	}

	if isDiffString(oldItem.Notes, item.Notes) {
		itemUpdate.Notes.Scan(item.Notes)
		itemNotes = append(itemNotes, model.UpdateItemNote{
			ValueTitle: model.UpdateableFieldNotes,
			ValueOld:   oldItem.Notes,
			ValueNew:   item.Notes,
		})
	}

	return itemUpdate, itemNotes
}

func isDiffString(oldItem, item string) bool {
	if len(oldItem) == 0 {
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
