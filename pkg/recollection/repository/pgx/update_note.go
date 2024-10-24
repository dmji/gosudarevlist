package repository_pgx

import (
	"collector/pkg/recollection/model"
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
)

func (r *repository) InsertUpdateNote(ctx context.Context, params model.UpdateNote) error {

	pgxDate, err := timeToPgTimestamp(params.UpdateDate)
	if err != nil {
		return err
	}

	err = r.query.InsertUpdateNote(ctx, pgx_sqlc.InsertUpdateNoteParams{
		ItemID:      params.ItemID,
		UpdateDate:  pgxDate,
		UpdateTitle: params.UpdateTitle,
		ValueOld:    params.ValueOld,
		ValueNew:    params.ValueNew,
	})

	return err

}

func (r *repository) GetUpdateNote(ctx context.Context, params model.OptionsGetNotes) ([]model.UpdateNote, error) {

	items, err := r.query.GetUpdateNote(ctx, pgx_sqlc.GetUpdateNoteParams{
		OffsetCount: params.Offset,
		Count:       params.Count,
	})

	if err != nil {
		return nil, err
	}

	res := make([]model.UpdateNote, 0, len(items))
	for _, item := range items {
		res = append(res, model.UpdateNote{
			ItemID:      item.ItemID,
			UpdateDate:  timeFromPgTimestamp(item.UpdateDate),
			UpdateTitle: item.Title,
			ValueOld:    item.ValueOld,
			ValueNew:    item.ValueNew,
		})
	}

	return res, nil

}
