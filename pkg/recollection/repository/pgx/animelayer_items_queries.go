package animelayer_repository

/*
import (
	animelayer_model "collector/pkg/animelayer/model"
	sqlc "collector/pkg/animelayer/repository/sqlc"
	"context"
)

func (r *repository) GetItemByIdentifier(ctx context.Context, identifier string) (*animelayer.ItemPartial, error) {

	item, err := r.query.GetItemByIdentifier(ctx, identifier)
	if err != nil {
		return nil, err
	}

	res := &animelayer.ItemPartial{
		Identifier:  item.Identifier,
		Title:       item.Title,
		IsCompleted: item.IsCompleted,
	}

	return res, nil
}

func (r *repository) InsertItem(ctx context.Context, item *animelayer.ItemPartial) error {

	return r.query.InsertNewItem(ctx,
		sqlc.InsertNewItemParams{
			Identifier:  item.Identifier,
			Title:       item.Title,
			IsCompleted: item.IsCompleted,
		},
	)

}

func (r *repository) UpdateItem(ctx context.Context, item *animelayer.ItemPartial) error {

	return r.query.UpdateItem(ctx,
		sqlc.UpdateItemParams{
			Identifier:  item.Identifier,
			Title:       item.Title,
			IsCompleted: item.IsCompleted,
		},
	)

}
*/
