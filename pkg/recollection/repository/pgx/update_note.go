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
