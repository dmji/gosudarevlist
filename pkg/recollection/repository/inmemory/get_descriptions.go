package repository_inmemory

import (
	animelayer_model "collector/pkg/animelayer/model"
	"context"
	"fmt"
)

func (r *repository) GetDescription(ctx context.Context, guid string) (animelayer_model.Description, error) {
	//log.Printf("In-Memory repo | GetDescriptions guid: %v", guid)

	for _, d := range r.descriptions {
		if d.Identifier == guid {
			return d, nil
		}

	}

	return animelayer_model.Description{}, fmt.Errorf("not found description for guid: %v", guid)
}
