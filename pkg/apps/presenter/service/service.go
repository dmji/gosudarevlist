package service

import (
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/repository"
)

type service struct {
	repository.AnimeLayerRepositoryDriver
}

func New(repo repository.AnimeLayerRepositoryDriver) *service {
	return &service{
		AnimeLayerRepositoryDriver: repo,
	}
}
