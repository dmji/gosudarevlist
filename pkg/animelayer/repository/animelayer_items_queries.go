package animelayer_repository

import (
	animelayer_model "collector/pkg/animelayer/model"
	sqlc "collector/pkg/animelayer/repository/sqlc"
	"context"
)

func (r repository) GetItemByIdentifier(ctx context.Context, identifier string) (*animelayer_model.Item, error) {

	item, err := r.query.GetItemByIdentifier(ctx, identifier)
	if err != nil {
		return nil, err
	}

	res := &animelayer_model.Item{
		GUID:      item.Identifier,
		Name:      item.Title,
		Completed: item.IsCompleted,
	}

	return res, nil
}

func (r repository) InsertItem(ctx context.Context, item *animelayer_model.Item) error {

	return r.query.InsertNewItem(ctx,
		sqlc.InsertNewItemParams{
			Identifier:  item.GUID,
			Title:       item.Name,
			IsCompleted: item.Completed,
		},
	)

}

func (r repository) UpdateItem(ctx context.Context, item *animelayer_model.Item) error {

	return r.query.UpdateItem(ctx,
		sqlc.UpdateItemParams{
			Identifier:  item.GUID,
			Title:       item.Name,
			IsCompleted: item.Completed,
		},
	)

}
