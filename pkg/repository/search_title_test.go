package repository_test

import (
	"collector/pkg/repository"
	"context"
	"fmt"
	"testing"
)

func TestSearchTitle(t *testing.T) {
	repo := repository.New()

	ctx := context.Background()
	list, _ := repo.SearchTitle(ctx, "Атака")
	for i, l := range list {
		fmt.Printf("%d: %s\n", i, l.Name)
	}
}
