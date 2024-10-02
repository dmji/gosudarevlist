package repository_inmemory

import (
	"context"
	"fmt"

	"github.com/dmji/go-animelayer-parser"
)

func (r *repository) GetDescription(ctx context.Context, guid string) (animelayer.ItemDetailed, error) {
	//log.Printf("In-Memory repo | GetDescriptions guid: %v", guid)

	for _, d := range r.descriptions {
		if d.Identifier == guid {
			return d, nil
		}

	}

	return animelayer.ItemDetailed{}, fmt.Errorf("not found description for guid: %v", guid)
}
