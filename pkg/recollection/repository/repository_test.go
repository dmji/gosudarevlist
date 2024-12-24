package repository_test

import (
	"testing"

	"github.com/dmji/gosudarevlist/pkg/recollection/repository"
	repository_pgx "github.com/dmji/gosudarevlist/pkg/recollection/repository/pgx"
)

func TestInterfacePgx(t *testing.T) {
	repo := repository_pgx.New(nil)

	var face repository.AnimeLayerRepositoryDriver
	face = repo
	_ = face
}
