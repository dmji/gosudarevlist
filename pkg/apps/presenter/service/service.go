package service

import (
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/repository"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
}

func New(repo repository.AnimeLayerRepositoryDriver) *services {
	return &services{
		AnimeLayerRepositoryDriver: repo,
	}
}
