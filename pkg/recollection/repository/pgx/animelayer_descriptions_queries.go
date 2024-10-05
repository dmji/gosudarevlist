package repository_pgx

/*
import (
	animelayer_model "collector/pkg/animelayer/model"
	sqlc "collector/pkg/animelayer/repository/sqlc"
	animelayer_comparator "collector/pkg/comparator"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *repository) GetDescriptionById(ctx context.Context, id int32) (*animelayer.ItemDetailed, error) {

	description, err := r.query.GetDescriptionById(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &animelayer.ItemDetailed{
		//Identifier: ,

		TorrentFilesSize: description.TorrentFilesSize,

		RefImagePreview: description.RefImagePreview,
		RefImageCover:   description.RefImageCover,

		UpdatedDate: description.UpdatedDate,
		CreatedDate: description.CreatedDate,

		LastCheckedDate: description.LastCheckedDate.Time.Format("2006-01-02 15:04"),
	}

	notesDb, err := r.query.GetDescriptionNote(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, note := range notesDb {

		res.Notes = append(res.Notes, animelayer_model.DescriptionNote{
			Name: note.FieldName,
			Text: note.FieldText,
		})

	}

	return res, nil
}

func (r *repository) InsertDescription(ctx context.Context, description *animelayer.ItemDetailed) error {

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	q := r.query.WithTx(tx)

	//
	// Insert
	//
	lastCheckedDate := pgtype.Date{}
	err = lastCheckedDate.Scan(description.LastCheckedDate)
	if err != nil {
		return err
	}

	descriptionId, err := q.InsertNewDescription(ctx,
		sqlc.InsertNewDescriptionParams{
			FirstCheckedDate: lastCheckedDate,
			CreatedDate:      description.CreatedDate,
			UpdatedDate:      description.UpdatedDate,
			RefImageCover:    description.RefImageCover,
			RefImagePreview:  description.RefImagePreview,
			TorrentFilesSize: description.TorrentFilesSize,
		},
	)

	if err != nil {
		return err
	}

	//
	// Update Notes
	//
	for _, note := range description.Notes {

		err = q.InsertNewDescriptionNote(ctx,
			sqlc.InsertNewDescriptionNoteParams{
				DescriptionID: descriptionId,
				FieldName:     note.Name,
				FieldText:     note.Text,
			},
		)
		if err != nil {
			return err
		}

	}

	//
	// Update Item
	//
	err = q.UpdateItemDescriptionId(ctx,
		sqlc.UpdateItemDescriptionIdParams{
			Identifier:    description.Identifier,
			DescriptionID: descriptionId,
		},
	)
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

func (r *repository) UpdateDescription(ctx context.Context, description *animelayer.ItemDetailed) error {

	lastCheckedDate := pgtype.Date{}
	err := lastCheckedDate.Scan(description.LastCheckedDate)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	q := r.query.WithTx(tx)

	//
	// Get Item
	//
	item, err := q.GetItemByIdentifier(ctx, description.Identifier)
	if err != nil {
		return err
	}

	descriptionID := item.DescriptionID

	//
	// Get Existed Description
	//
	descriptionCurrent, err := r.GetDescriptionById(ctx, descriptionID)
	if err != nil {
		return err
	}

	difference, notesUpdateIndexes := animelayer_comparator.CompareDescriptions(description, descriptionCurrent)
	for _, d := range difference {
		q.InsertNewUpdateNote(ctx, sqlc.InsertNewUpdateNoteParams{
			Item:       item.ItemID,
			UpdateDate: lastCheckedDate,
			Title:      d.Name,
			ValueOld:   d.OldValue,
			ValueNew:   d.NewValue,
		})
	}

	//
	// Update Description
	//

	err = q.UpdateDescription(ctx,
		sqlc.UpdateDescriptionParams{
			LastCheckedDate:  lastCheckedDate,
			CreatedDate:      description.CreatedDate,
			UpdatedDate:      description.UpdatedDate,
			RefImageCover:    description.RefImageCover,
			RefImagePreview:  description.RefImagePreview,
			TorrentFilesSize: description.TorrentFilesSize,
			DescriptionID:    descriptionID,
		},
	)
	if err != nil {
		return err
	}

	//
	// Update Description Notes
	//
	for i := range notesUpdateIndexes.IndexOldRemoved {
		q.DeleteDescriptionNote(ctx,
			sqlc.DeleteDescriptionNoteParams{
				DescriptionID: descriptionID,
				FieldName:     descriptionCurrent.Notes[i].Name,
			},
		)

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
*/
