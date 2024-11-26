package repository_test

import (
	"collector/pkg/recollection/repository"
	repository_pgx "collector/pkg/recollection/repository/pgx"
	"testing"
)

func TestInterfacePgx(t *testing.T) {
	repo := repository_pgx.New(nil)

	var face repository.AnimeLayerRepositoryDriver
	face = repo
	_ = face
}
