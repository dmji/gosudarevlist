package service

import (
	"bytes"
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/dmji/gosudarevlist/internal/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/websocket"
)

type service struct {
	repo repository.AnimeLayerRepositoryDriver
	data sync.Map // map[enums.Category]*categoryUpdaterData
}

func New(repo repository.AnimeLayerRepositoryDriver) *service {
	s := &service{
		repo: repo,
	}

	return s
}

func (s *service) UpdateTrigger(ctx context.Context, cat enums.Category) {
	data := s.updaterDataByCategory(ctx, cat)
	if data == nil {
		logger.Errorw(ctx, "Update Trigger Data not found")
		return
	}

	data.ws.PublishTempl(data.publishUpdate())
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
		data.ws = websocket.NewManager(
			category.String()+" Updater",
			10,
			userDataInitializer,
			func(ctx context.Context, _ *WsUserData) []byte {
				buf := &bytes.Buffer{}
				data.publishUpdate()(ctx, buf)
				return buf.Bytes()
			},
		)

		s.data.Store(category, data)
		return data
	}

	return dataPtr.(*categoryUpdaterData)
}

func (s *service) SubscribeHandler(ctx context.Context, category enums.Category) func(w http.ResponseWriter, r *http.Request) {
	return s.updaterDataByCategory(ctx, category).ws.SubscribeHandler
}
