package service

import (
	"context"
	"sync"
	"time"

	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

type ItemProvider interface {
	GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error)
	GetItemsFromCategoryPages(ctx context.Context, category enums.Category, iPage int) ([]*model.AnimelayerItem, error)
}

type UpdaterManagerNotifier interface {
	UpdateTrigger(ctx context.Context, cat enums.Category)
}

type service struct {
	repo            repository.AnimeLayerRepositoryDriver
	animelayerApi   ItemProvider
	managerNotifier UpdaterManagerNotifier

	data sync.Map // map[enums.Category]*categoryUpdaterData
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerApi ItemProvider, managerNotifier UpdaterManagerNotifier) *service {
	s := &service{
		repo:            repo,
		animelayerApi:   animelayerApi,
		managerNotifier: managerNotifier,
	}

	return s
}

type categoryUpdaterData struct {
	lastUpdateTimer time.Time
	mx              sync.Mutex
	category        enums.Category
}

func (s *service) updaterDataByCategory(ctx context.Context, category enums.Category) *categoryUpdaterData {
	dataPtr, ok := s.data.Load(category)
	if !ok {

		timeLastUpdate, err := s.repo.GetLastCategoryUpdateItem(ctx, category)
		if err != nil {
			t := time.Now().Add(-10 * time.Second)
			timeLastUpdate = &t
		}

		data := &categoryUpdaterData{
			lastUpdateTimer: *timeLastUpdate,
			category:        category,
		}

		s.data.Store(category, data)
		return data
	}

	return dataPtr.(*categoryUpdaterData)
}
