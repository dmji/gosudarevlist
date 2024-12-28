package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func (r *repository) GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error) {

	item, err := r.query.GetItemByIdentifier(ctx, identifier)
	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItemByIdentifier", "error", err)
		return nil, err
	}

	return &model.AnimelayerItem{
		Identifier:       identifier,
		Title:            item.Title,
		IsCompleted:      item.IsCompleted,
		LastCheckedDate:  timeFromPgTimestamp(item.LastCheckedDate),
		FirstCheckedDate: timeFromPgTimestamp(item.FirstCheckedDate),
		CreatedDate:      timeFromPgTimestamp(item.CreatedDate),
		UpdatedDate:      timeFromPgTimestamp(item.UpdatedDate),
		RefImageCover:    item.RefImageCover,
		RefImagePreview:  item.RefImagePreview,
		BlobImageCover:   item.BlobImageCover,
		BlobImagePreview: item.BlobImagePreview,
		TorrentFilesSize: item.TorrentFilesSize,
		Notes:            item.Notes,
		Category:         pgxCategoriesToCategory(item.Category),
	}, nil
}
