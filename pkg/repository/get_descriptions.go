package repository

import (
	"collector/pkg/model"
	"context"
	"fmt"
)

func (r *repository) GetDescription(ctx context.Context, guid string) (model.AnimeLayerItemDescription, error) {
	//log.Printf("In-Memory repo | GetDescriptions guid: %v", guid)

	for _, d := range r.descriptions {
		if d.GUID == guid {
			return d, nil
		}

	}

	return model.AnimeLayerItemDescription{}, fmt.Errorf("not found description for guid: %v", guid)
}
