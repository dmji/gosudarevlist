package animelayer_repository

import (
	animelayer_model "collector/pkg/animelayer/model"
	sqlc "collector/pkg/animelayer/repository/sqlc"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

/*
	 func (r repository) GetDescriptionByIdentifier(ctx context.Context, identifier string) (*animelayer_model.Description, error) {

		item, err := r.query.GetDescriptionByIdentifier(ctx, identifier)
		if err != nil {
			return nil, err
		}

		res := &animelayer_model.Description{
			Identifier: item.Identifier,
			Name:       item.Title,
			Completed:  item.IsCompleted,
		}

		return res, nil
	}
*/
func (r repository) InsertDescription(ctx context.Context, item *animelayer_model.Description) error {

	lastCheckedDate := pgtype.Date{}
	err := lastCheckedDate.Scan(item.LastCheckedDate)
	if err != nil {
		return err
	}

	return r.query.InsertNewDescription(ctx,
		sqlc.InsertNewDescriptionParams{
			FirstCheckedDate: lastCheckedDate,
			CreatedDate:      item.CreatedDate,
			UpdatedDate:      item.UpdatedDate,
			RefImageCover:    item.RefImageCover,
			RefImagePreview:  item.RefImagePreview,
			TorrentFilesSize: item.TorrentFilesSize,
		},
	)

}

func (r repository) UpdateDescription(ctx context.Context, item *animelayer_model.Description) error {

	lastCheckedDate := pgtype.Date{}
	err := lastCheckedDate.Scan(item.LastCheckedDate)
	if err != nil {
		return err
	}

	return r.query.UpdateDescription(ctx,
		sqlc.UpdateDescriptionParams{
			LastCheckedDate:  lastCheckedDate,
			CreatedDate:      item.CreatedDate,
			UpdatedDate:      item.UpdatedDate,
			RefImageCover:    item.RefImageCover,
			RefImagePreview:  item.RefImagePreview,
			TorrentFilesSize: item.TorrentFilesSize,
			DescriptionID:    1,
		},
	)
}
