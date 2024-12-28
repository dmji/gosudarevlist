package repository_pgx

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func (r *repository) InsertUpdateNote(ctx context.Context, params model.UpdateItem) error {

	/* 	pgxDate, err := timeToPgTimestamp(params.Date)
	   	if err != nil {
	   		return err
	   	} */

	/* 	err = r.query.InsertUpdateNote(ctx, pgx_sqlc.InsertUpdateNoteParams{
		ItemID:      params.ItemID,
		UpdateDate:  pgxDate,
		UpdateTitle: params.Title,
		ValueOld:    params.ValueOld,
		ValueNew:    params.ValueNew,
	})
	if err != nil {
	   		return err
	   	}
	*/

	return nil

}
